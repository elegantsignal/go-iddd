package domain_test

import (
	"go-iddd/customer/domain"
	"go-iddd/customer/domain/commands"
	"go-iddd/customer/domain/events"
	"go-iddd/customer/domain/values"
	"go-iddd/shared"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestChangeEmailAddressOfCustomer(t *testing.T) {
	Convey("Given a Customer", t, func() {
		id, err := values.CustomerIDFrom("64bcf656-da30-4f5a-b0b5-aead60965aa3")
		So(err, ShouldBeNil)
		emailAddress, err := values.EmailAddressFrom("john@doe.com")
		So(err, ShouldBeNil)
		confirmableEmailAddress := emailAddress.ToConfirmable()
		personName, err := values.PersonNameFrom("John", "Doe")
		So(err, ShouldBeNil)

		currentStreamVersion := uint(1)

		customer, err := domain.ReconstituteCustomerFrom(
			shared.DomainEvents{
				events.ItWasRegistered(id, confirmableEmailAddress, personName, currentStreamVersion),
			},
		)
		So(err, ShouldBeNil)

		Convey("When an emailAddress is changed", func() {
			newEmailAddress, err := values.EmailAddressFrom("john+changed@doe.com")
			So(err, ShouldBeNil)

			changeEmailAddress, err := commands.NewChangeEmailAddress(
				id.String(),
				newEmailAddress.EmailAddress(),
			)
			So(err, ShouldBeNil)

			recordedEvents := domain.ChangeEmailAddress(customer, changeEmailAddress)

			Convey("It should record EmailAddressChanged", func() {
				So(recordedEvents, ShouldHaveLength, 1)
				emailAddressChanged, ok := recordedEvents[0].(*events.EmailAddressChanged)
				So(ok, ShouldBeTrue)
				So(emailAddressChanged, ShouldNotBeNil)
				So(emailAddressChanged.CustomerID().Equals(id), ShouldBeTrue)
				So(emailAddressChanged.ConfirmableEmailAddress().Equals(newEmailAddress), ShouldBeTrue)
				So(emailAddressChanged.StreamVersion(), ShouldEqual, currentStreamVersion+1)

				Convey("And when it is changed to the same value again", func() {
					recordedEvents := domain.ChangeEmailAddress(customer, changeEmailAddress)

					Convey("It should be ignored", func() {
						So(recordedEvents, ShouldBeEmpty)
					})
				})
			})
		})
	})
}
