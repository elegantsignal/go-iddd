package grpc

import (
	"database/sql"

	"github.com/AntonStoeckl/go-iddd/src/customeraccounts/hexagon/application"
	"github.com/AntonStoeckl/go-iddd/src/customeraccounts/hexagon/application/domain/customer"
	customergrpc "github.com/AntonStoeckl/go-iddd/src/customeraccounts/infrastructure/adapter/grpc"
	customergrpcproto "github.com/AntonStoeckl/go-iddd/src/customeraccounts/infrastructure/adapter/grpc/proto"
	"github.com/AntonStoeckl/go-iddd/src/customeraccounts/infrastructure/adapter/postgres"
	"github.com/AntonStoeckl/go-iddd/src/customeraccounts/infrastructure/serialization"
	"github.com/AntonStoeckl/go-iddd/src/shared"
	"github.com/AntonStoeckl/go-iddd/src/shared/es"
	"github.com/cockroachdb/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	eventStoreTableName           = "eventstore"
	uniqueEmailAddressesTableName = "unique_email_addresses"
)

type DIOption func(container *DIContainer) error

func UsePostgresDBConn(dbConn *sql.DB) DIOption {
	return func(container *DIContainer) error {
		if dbConn == nil {
			return errors.New("pgDBConn must not be nil")
		}

		container.infra.pgDBConn = dbConn

		return nil
	}
}

func WithMarshalCustomerEvents(fn es.MarshalDomainEvent) DIOption {
	return func(container *DIContainer) error {
		container.dependency.marshalCustomerEvent = fn
		return nil
	}
}

func WithUnmarshalCustomerEvents(fn es.UnmarshalDomainEvent) DIOption {
	return func(container *DIContainer) error {
		container.dependency.unmarshalCustomerEvent = fn
		return nil
	}
}

func WithBuildUniqueEmailAddressAssertions(fn customer.ForBuildingUniqueEmailAddressAssertions) DIOption {
	return func(container *DIContainer) error {
		container.dependency.buildUniqueEmailAddressAssertions = fn
		return nil
	}
}

func ReplaceGRPCCustomerServer(server customergrpcproto.CustomerServer) DIOption {
	return func(container *DIContainer) error {
		if server == nil {
			return errors.New("grpcCustomerServer must not be nil")
		}

		container.service.grpcCustomerServer = server

		return nil
	}
}

type DIContainer struct {
	config *Config

	infra struct {
		pgDBConn *sql.DB
	}

	dependency struct {
		marshalCustomerEvent              es.MarshalDomainEvent
		unmarshalCustomerEvent            es.UnmarshalDomainEvent
		buildUniqueEmailAddressAssertions customer.ForBuildingUniqueEmailAddressAssertions
	}

	service struct {
		customerEventStore     *postgres.CustomerEventStore
		customerCommandHandler *application.CustomerCommandHandler
		customerQueryHandler   *application.CustomerQueryHandler
		grpcCustomerServer     customergrpcproto.CustomerServer
		grpcServer             *grpc.Server
	}
}

func MustBuildDIContainer(config *Config, logger *shared.Logger, opts ...DIOption) *DIContainer {
	container := &DIContainer{}
	container.config = config

	/*** Define default dependencies ***/
	container.dependency.marshalCustomerEvent = serialization.MarshalCustomerEvent
	container.dependency.unmarshalCustomerEvent = serialization.UnmarshalCustomerEvent
	container.dependency.buildUniqueEmailAddressAssertions = customer.BuildUniqueEmailAddressAssertions

	/*** Apply options for infra, dependencies, services ***/
	for _, opt := range opts {
		if err := opt(container); err != nil {
			logger.Panic().Msgf("mustBuildDIContainer: %s", err)
		}
	}

	container.init()

	return container
}

func (container *DIContainer) init() {
	_ = container.GetCustomerEventStore()
	_ = container.GetCustomerCommandHandler()
	_ = container.GetCustomerQueryHandler()
	_ = container.GetGRPCCustomerServer()
	_ = container.GetGRPCServer()
}

func (container *DIContainer) GetPostgresDBConn() *sql.DB {
	return container.infra.pgDBConn
}

func (container *DIContainer) GetCustomerEventStore() *postgres.CustomerEventStore {
	if container.service.customerEventStore == nil {
		container.service.customerEventStore = postgres.NewCustomerEventStore(
			container.infra.pgDBConn,
			eventStoreTableName,
			container.dependency.marshalCustomerEvent,
			container.dependency.unmarshalCustomerEvent,
			uniqueEmailAddressesTableName,
			container.dependency.buildUniqueEmailAddressAssertions,
		)
	}

	return container.service.customerEventStore
}

func (container *DIContainer) GetCustomerCommandHandler() *application.CustomerCommandHandler {
	if container.service.customerCommandHandler == nil {
		container.service.customerCommandHandler = application.NewCustomerCommandHandler(
			container.GetCustomerEventStore().RetrieveEventStream,
			container.GetCustomerEventStore().StartEventStream,
			container.GetCustomerEventStore().AppendToEventStream,
		)
	}

	return container.service.customerCommandHandler
}

func (container *DIContainer) GetCustomerQueryHandler() *application.CustomerQueryHandler {
	if container.service.customerQueryHandler == nil {
		container.service.customerQueryHandler = application.NewCustomerQueryHandler(
			container.GetCustomerEventStore().RetrieveEventStream,
		)
	}

	return container.service.customerQueryHandler
}

func (container *DIContainer) GetGRPCCustomerServer() customergrpcproto.CustomerServer {
	if container.service.grpcCustomerServer == nil {
		container.service.grpcCustomerServer = customergrpc.NewCustomerServer(
			container.GetCustomerCommandHandler().RegisterCustomer,
			container.GetCustomerCommandHandler().ConfirmCustomerEmailAddress,
			container.GetCustomerCommandHandler().ChangeCustomerEmailAddress,
			container.GetCustomerCommandHandler().ChangeCustomerName,
			container.GetCustomerCommandHandler().DeleteCustomer,
			container.GetCustomerQueryHandler().CustomerViewByID,
		)
	}

	return container.service.grpcCustomerServer
}

func (container *DIContainer) GetGRPCServer() *grpc.Server {
	if container.service.grpcServer == nil {
		container.service.grpcServer = grpc.NewServer()
		customergrpcproto.RegisterCustomerServer(container.service.grpcServer, container.GetGRPCCustomerServer())
		reflection.Register(container.service.grpcServer)
	}

	return container.service.grpcServer
}
