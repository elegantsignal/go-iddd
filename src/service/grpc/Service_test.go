package grpc_test

import (
	"context"
	"fmt"
	"syscall"
	"testing"
	"time"

	"github.com/AntonStoeckl/go-iddd/src/customeraccounts/hexagon/application/domain/customer"
	"github.com/AntonStoeckl/go-iddd/src/customeraccounts/hexagon/application/domain/customer/value"
	customergrpc "github.com/AntonStoeckl/go-iddd/src/customeraccounts/infrastructure/adapter/grpc"
	customergrpcproto "github.com/AntonStoeckl/go-iddd/src/customeraccounts/infrastructure/adapter/grpc/proto"
	grpcService "github.com/AntonStoeckl/go-iddd/src/service/grpc"
	"github.com/AntonStoeckl/go-iddd/src/shared"
	. "github.com/smartystreets/goconvey/convey"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestStartGRPCServer(t *testing.T) {
	logger := shared.NewNilLogger()
	config := grpcService.MustBuildConfigFromEnv(logger)
	postgresDBConn := grpcService.MustInitPostgresDB(config, logger)
	diContainer := grpcService.MustBuildDIContainer(
		config,
		logger,
		grpcService.UsePostgresDBConn(postgresDBConn),
		grpcService.ReplaceGRPCCustomerServer(grpcCustomerServerStub()),
	)

	exitWasCalled := false
	exitFn := func() {
		exitWasCalled = true
	}

	terminateDelay := time.Millisecond * 100

	s := grpcService.InitService(config, logger, exitFn, diContainer)

	Convey("Start the gRPC server as a goroutine", t, func() {
		go s.StartGRPCServer()

		Convey("gPRC server should handle requests", func() {
			client := customerGRPCClient(config)
			res, err := client.Register(context.Background(), &customergrpcproto.RegisterRequest{})
			So(err, ShouldBeNil)
			So(res, ShouldNotBeNil)
			So(res.Id, ShouldNotBeEmpty)

			Convey(fmt.Sprintf("It should wait for stop signal (scheduled after %s)", terminateDelay), func() {
				start := time.Now()
				go func() {
					time.Sleep(terminateDelay)
					_ = syscall.Kill(syscall.Getpid(), syscall.SIGTERM)
				}()

				s.WaitForStopSignal()

				So(time.Now(), ShouldHappenOnOrAfter, start.Add(terminateDelay))

				Convey("Stop signal should issue Shutdown", func() {
					Convey("Shutdown should stop gRPC server", func() {
						_, err = client.Register(context.Background(), &customergrpcproto.RegisterRequest{})
						So(err, ShouldBeError)
						So(status.Code(err), ShouldResemble, codes.Unavailable)

						Convey("Shutdown should close PostgreSQL connection", func() {
							err := postgresDBConn.Ping()
							So(err, ShouldBeError)
							So(err.Error(), ShouldContainSubstring, "database is closed")

							Convey("Shutdown should call exit", func() {
								So(exitWasCalled, ShouldBeTrue)
							})
						})
					})
				})
			})
		})
	})
}

/*** Helper functions ***/

func grpcCustomerServerStub() customergrpcproto.CustomerServer {
	customerServer := customergrpc.NewCustomerServer(
		func(customerIDValue value.CustomerID, emailAddress, givenName, familyName string) error {
			return nil
		},
		func(customerID, confirmationHash string) error {
			return nil
		},
		func(customerID, emailAddress string) error {
			return nil
		},
		func(customerID, givenName, familyName string) error {
			return nil
		},
		func(customerID string) error {
			return nil
		},
		func(customerID string) (customer.View, error) {
			return customer.View{}, nil
		},
	)

	return customerServer
}

func customerGRPCClient(config *grpcService.Config) customergrpcproto.CustomerClient {
	grpcClientConn, _ := grpc.DialContext(context.Background(), config.GRPC.HostAndPort, grpc.WithInsecure(), grpc.WithBlock())
	client := customergrpcproto.NewCustomerClient(grpcClientConn)

	return client
}
