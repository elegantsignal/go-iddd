package commands_test

import (
	"go-iddd/customer/domain/commands"
	"go-iddd/customer/domain/values"
	"go-iddd/shared"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"golang.org/x/xerrors"
)

func TestNewConfirmEmailAddress(t *testing.T) {
	Convey("Given valid ID, EmailAddress and ConfirmationHash", t, func() {
		id := "64bcf656-da30-4f5a-b0b5-aead60965aa3"
		emailAddress := "foo@bar.com"
		confirmationHash := "secret_hash"

		Convey("When a new ConfirmEmailAddress command is created", func() {
			confirmEmailAddress, err := commands.NewConfirmEmailAddress(
				id,
				emailAddress,
				confirmationHash,
			)

			Convey("It should succeed", func() {
				So(err, ShouldBeNil)
				So(confirmEmailAddress, ShouldHaveSameTypeAs, (*commands.ConfirmEmailAddress)(nil))
			})
		})

		Convey("Given that ID is invalid instead", func() {
			id = ""
			conveyNewConfirmEmailAddressWithInvalidInput(id, emailAddress, confirmationHash)
		})

		Convey("Given that EmailAddress is invalid instead", func() {
			emailAddress = ""
			conveyNewConfirmEmailAddressWithInvalidInput(id, emailAddress, confirmationHash)
		})

		Convey("Given that ConfirmationHash is invalid instead", func() {
			confirmationHash = ""
			conveyNewConfirmEmailAddressWithInvalidInput(id, emailAddress, confirmationHash)
		})
	})
}

func conveyNewConfirmEmailAddressWithInvalidInput(
	id string,
	emailAddress string,
	confirmationHash string,
) {

	Convey("When a new ConfirmEmailAddress command is created", func() {
		confirmEmailAddress, err := commands.NewConfirmEmailAddress(
			id,
			emailAddress,
			confirmationHash,
		)

		Convey("It should fail", func() {
			So(err, ShouldBeError)
			So(xerrors.Is(err, shared.ErrInputIsInvalid), ShouldBeTrue)
			So(confirmEmailAddress, ShouldBeNil)
		})
	})
}

func TestConfirmEmailAddressExposesExpectedValues(t *testing.T) {
	Convey("Given a ConfirmEmailAddress command", t, func() {
		id := "64bcf656-da30-4f5a-b0b5-aead60965aa3"
		emailAddress := "foo@bar.com"
		confirmationHash := "secret_hash"

		idValue, err := values.NewID(id)
		So(err, ShouldBeNil)
		emailAddressValue, err := values.NewEmailAddress("foo@bar.com")
		So(err, ShouldBeNil)
		confirmationHashValue, err := values.NewConfirmationHash(confirmationHash)
		So(err, ShouldBeNil)

		confirmEmailAddress, err := commands.NewConfirmEmailAddress(
			id,
			emailAddress,
			confirmationHash,
		)
		So(err, ShouldBeNil)

		Convey("It should expose the expected values", func() {
			So(idValue.Equals(confirmEmailAddress.ID()), ShouldBeTrue)
			So(emailAddressValue.Equals(confirmEmailAddress.EmailAddress()), ShouldBeTrue)
			So(confirmationHashValue.Equals(confirmEmailAddress.ConfirmationHash()), ShouldBeTrue)
			So(confirmEmailAddress.CommandName(), ShouldEqual, "ConfirmEmailAddress")
			So(idValue.Equals(confirmEmailAddress.AggregateIdentifier()), ShouldBeTrue)
		})
	})
}
