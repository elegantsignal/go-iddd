package customergrpc_test

import (
	"context"
	"testing"

	"github.com/AntonStoeckl/go-iddd/src/customeraccounts/hexagon/application/domain/customer"
	"github.com/AntonStoeckl/go-iddd/src/customeraccounts/hexagon/application/domain/customer/value"
	customergrpc "github.com/AntonStoeckl/go-iddd/src/customeraccounts/infrastructure/adapter/grpc"
	customergrpcproto "github.com/AntonStoeckl/go-iddd/src/customeraccounts/infrastructure/adapter/grpc/proto"
	"github.com/AntonStoeckl/go-iddd/src/shared"
	"github.com/cockroachdb/errors"
	"github.com/golang/protobuf/ptypes/empty"
	. "github.com/smartystreets/goconvey/convey"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var generatedID value.CustomerID
var mockedView = customer.View{
	ID:                      generatedID.String(),
	EmailAddress:            "fiona@gallagher.net",
	IsEmailAddressConfirmed: true,
	GivenName:               "Fiona",
	FamilyName:              "Gallagher",
	IsDeleted:               false,
	Version:                 2,
}
var expectedErrCode = codes.InvalidArgument
var expectedErrMsg = "invalid input"

func TestGRPCServer(t *testing.T) {
	Convey("Prepare test artifacts", t, func() {
		successCustomerServer := buildSuccessCustomerServer()
		failureCustomerServer := buildFailureCustomerServer()

		Convey("Usecase: Register", func() {
			Convey("Given the application will return success", func() {
				Convey("When the request is handled", func() {
					res, err := successCustomerServer.Register(
						context.Background(),
						&customergrpcproto.RegisterRequest{},
					)

					Convey("Then it should succeed", func() {
						So(err, ShouldBeNil)
						So(res, ShouldNotBeNil)
						So(res.Id, ShouldResemble, generatedID.String())
					})
				})
			})

			Convey("Given the application will return an error", func() {
				Convey("When the request is handled", func() {
					res, err := failureCustomerServer.Register(
						context.Background(),
						&customergrpcproto.RegisterRequest{},
					)

					Convey("Then it should fail with the exptected error", func() {
						So(err, ShouldBeError)
						So(err, ShouldResemble, status.Error(expectedErrCode, expectedErrMsg))
						So(res, ShouldBeNil)
					})
				})
			})
		})

		Convey("\nUsecase: ConfirmEmailAddress", func() {
			Convey("Given the application will return success", func() {
				Convey("When the request is handled", func() {
					res, err := successCustomerServer.ConfirmEmailAddress(
						context.Background(),
						&customergrpcproto.ConfirmEmailAddressRequest{},
					)

					thenItShouldSuccees(res, err)
				})
			})

			Convey("Given the application will return an error", func() {
				Convey("When the request is handled", func() {
					res, err := failureCustomerServer.ConfirmEmailAddress(
						context.Background(),
						&customergrpcproto.ConfirmEmailAddressRequest{},
					)

					thenItShouldFailWithTheExpectedError(res, err)
				})
			})
		})

		Convey("\nUsecase: ChangeEmailAddress", func() {
			Convey("Given the application will return success", func() {
				Convey("When the request is handled", func() {
					res, err := successCustomerServer.ChangeEmailAddress(
						context.Background(),
						&customergrpcproto.ChangeEmailAddressRequest{},
					)

					thenItShouldSuccees(res, err)
				})
			})

			Convey("Given the application will return an error", func() {
				Convey("When the request is handled", func() {
					res, err := failureCustomerServer.ChangeEmailAddress(
						context.Background(),
						&customergrpcproto.ChangeEmailAddressRequest{},
					)

					thenItShouldFailWithTheExpectedError(res, err)
				})
			})
		})

		Convey("\nUsecase: ChangeName", func() {
			Convey("Given the application will return success", func() {
				Convey("When the request is handled", func() {
					res, err := successCustomerServer.ChangeName(
						context.Background(),
						&customergrpcproto.ChangeNameRequest{},
					)

					thenItShouldSuccees(res, err)
				})
			})

			Convey("Given the application will return an error", func() {
				Convey("When the request is handled", func() {
					res, err := failureCustomerServer.ChangeName(
						context.Background(),
						&customergrpcproto.ChangeNameRequest{},
					)

					thenItShouldFailWithTheExpectedError(res, err)
				})
			})
		})

		Convey("\nUsecase: Delete", func() {
			Convey("Given the application will return success", func() {
				Convey("When the request is handled", func() {
					res, err := successCustomerServer.Delete(
						context.Background(),
						&customergrpcproto.DeleteRequest{},
					)

					thenItShouldSuccees(res, err)
				})
			})

			Convey("Given the application will return an error", func() {
				Convey("When the request is handled", func() {
					res, err := failureCustomerServer.Delete(
						context.Background(),
						&customergrpcproto.DeleteRequest{},
					)

					thenItShouldFailWithTheExpectedError(res, err)
				})
			})
		})

		Convey("\nUsecase: RetrieveView", func() {
			Convey("Given the application will return success", func() {
				Convey("When the request is handled", func() {
					res, err := successCustomerServer.RetrieveView(
						context.Background(),
						&customergrpcproto.RetrieveViewRequest{},
					)

					Convey("Then it should succeed", func() {
						So(err, ShouldBeNil)
						So(res, ShouldNotBeNil)

						expectedRes := &customergrpcproto.RetrieveViewResponse{
							EmailAddress:            mockedView.EmailAddress,
							IsEmailAddressConfirmed: mockedView.IsEmailAddressConfirmed,
							GivenName:               mockedView.GivenName,
							FamilyName:              mockedView.FamilyName,
							Version:                 uint64(mockedView.Version),
						}

						So(res, ShouldResemble, expectedRes)
					})
				})
			})

			Convey("Given the application will return an error", func() {
				Convey("When the request is handled", func() {
					res, err := failureCustomerServer.RetrieveView(
						context.Background(),
						&customergrpcproto.RetrieveViewRequest{},
					)

					Convey("Then it should fail with the exptected error", func() {
						So(err, ShouldBeError)
						So(err, ShouldResemble, status.Error(expectedErrCode, expectedErrMsg))
						So(res, ShouldBeNil)
					})
				})
			})
		})
	})
}

func thenItShouldSuccees(res *empty.Empty, err error) {
	Convey("Then it should succeed", func() {
		So(err, ShouldBeNil)
		So(res, ShouldResemble, &empty.Empty{})
	})
}

func thenItShouldFailWithTheExpectedError(res *empty.Empty, err error) {
	Convey("Then it should fail with the exptected error", func() {
		So(err, ShouldBeError)
		So(err, ShouldResemble, status.Error(expectedErrCode, expectedErrMsg))
		So(res, ShouldBeNil)
	})
}

func buildSuccessCustomerServer() customergrpcproto.CustomerServer {
	customerGRPCServer := customergrpc.NewCustomerServer(
		func(customerIDValue value.CustomerID, emailAddress, givenName, familyName string) error {
			generatedID = customerIDValue
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
			return mockedView, nil
		},
	)

	return customerGRPCServer
}

func buildFailureCustomerServer() customergrpcproto.CustomerServer {
	mockedView := customer.View{}
	mockedErr := errors.Mark(errors.New(expectedErrMsg), shared.ErrInputIsInvalid)

	customerGRPCServer := customergrpc.NewCustomerServer(
		func(customerIDValue value.CustomerID, emailAddress, givenName, familyName string) error {
			return mockedErr
		},
		func(customerID, confirmationHash string) error {
			return mockedErr
		},
		func(customerID, emailAddress string) error {
			return mockedErr
		},
		func(customerID, givenName, familyName string) error {
			return mockedErr
		},
		func(customerID string) error {
			return mockedErr
		},
		func(customerID string) (customer.View, error) {
			return mockedView, mockedErr
		},
	)

	return customerGRPCServer
}
