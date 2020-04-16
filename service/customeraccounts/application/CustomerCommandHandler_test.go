package application_test

import (
	"testing"

	"github.com/AntonStoeckl/go-iddd/service/customeraccounts/application"
	"github.com/AntonStoeckl/go-iddd/service/customeraccounts/application/domain"
	"github.com/AntonStoeckl/go-iddd/service/customeraccounts/application/domain/customer/value"
	"github.com/AntonStoeckl/go-iddd/service/shared"
	"github.com/AntonStoeckl/go-iddd/service/shared/es"
	"github.com/cockroachdb/errors"
	. "github.com/smartystreets/goconvey/convey"
)

type commandHandlerTestArtifacts struct {
	emailAddress       string
	givenName          string
	familyName         string
	newEmailAddress    string
	newGivenName       string
	newFamilyName      string
	customerID         value.CustomerID
	confirmationHash   value.ConfirmationHash
	customerRegistered domain.CustomerRegistered
}

func TestCustomerCommandHandler_TechnicalProblemsWithCustomerEventStore(t *testing.T) {
	Convey("Prepare test artifacts", t, func() {
		var err error
		ca := buildArtifactsForCommandHandlerTest()

		Convey("\nSCENARIO: Technical problems with the CustomerEventStore", func() {
			Convey("Given a registered Customer", func() {
				Convey("and assuming the recorded events can't be stored", func() {
					commandHandler := application.NewCustomerCommandHandler(
						func(id value.CustomerID) (es.EventStream, error) {
							return es.EventStream{ca.customerRegistered}, nil
						},
						func(recordedEvents es.RecordedEvents, id value.CustomerID) error {
							return nil
						},
						func(recordedEvents es.RecordedEvents, id value.CustomerID) error {
							return shared.ErrTechnical
						},
						shared.RetryOnConcurrencyConflict,
					)

					Convey("When he tries to confirm his email address", func() {
						err = commandHandler.ConfirmCustomerEmailAddress(ca.customerID.String(), ca.confirmationHash.String())

						Convey("Then he should receive an error", func() {
							So(err, ShouldBeError)
							So(errors.Is(err, shared.ErrTechnical), ShouldBeTrue)
						})
					})

					Convey("When he tries to change his email address", func() {
						err = commandHandler.ChangeCustomerEmailAddress(ca.customerID.String(), ca.newEmailAddress)

						Convey("Then he should receive an error", func() {
							So(err, ShouldBeError)
							So(errors.Is(err, shared.ErrTechnical), ShouldBeTrue)
						})
					})

					Convey("When he tries to change his name", func() {
						err = commandHandler.ChangeCustomerName(ca.customerID.String(), ca.givenName, ca.familyName)

						Convey("Then he should receive an error", func() {
							So(err, ShouldBeError)
							So(errors.Is(err, shared.ErrTechnical), ShouldBeTrue)
						})
					})

					Convey("When he tries to delete his account", func() {
						err = commandHandler.DeleteCustomer(ca.customerID.String())

						Convey("Then he should receive an error", func() {
							So(err, ShouldBeError)
							So(errors.Is(err, shared.ErrTechnical), ShouldBeTrue)
						})
					})
				})
			})
		})
	})
}

func buildArtifactsForCommandHandlerTest() commandHandlerTestArtifacts {
	ca := commandHandlerTestArtifacts{}

	ca.emailAddress = "fiona@gallagher.net"
	ca.givenName = "Fiona"
	ca.familyName = "Galagher"
	ca.newEmailAddress = "fiona@pratt.net"
	ca.newGivenName = "Fiona"
	ca.newFamilyName = "Pratt"

	ca.customerID = value.GenerateCustomerID()
	ca.confirmationHash = value.GenerateConfirmationHash(ca.emailAddress)

	ca.customerRegistered = domain.BuildCustomerRegistered(
		ca.customerID,
		value.RebuildEmailAddress(ca.emailAddress),
		ca.confirmationHash,
		value.RebuildPersonName(ca.givenName, ca.familyName),
		1,
	)

	return ca
}