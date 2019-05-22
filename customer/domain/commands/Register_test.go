package commands_test

import (
	"go-iddd/customer/domain/commands"
	"go-iddd/customer/domain/values"
	"go-iddd/shared"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"golang.org/x/xerrors"
)

func TestNewRegister(t *testing.T) {
	Convey("Given valid input", t, func() {
		id := "64bcf656-da30-4f5a-b0b5-aead60965aa3"
		emailAddress := "john@doe.com"
		givenName := "John"
		familyName := "Doe"

		Convey("When a new Register command is created", func() {
			register, err := commands.NewRegister(id, emailAddress, givenName, familyName)

			Convey("It should succeed", func() {
				So(err, ShouldBeNil)
				So(register, ShouldHaveSameTypeAs, (*commands.Register)(nil))
			})
		})

		Convey("Given that id is invalid", func() {
			id = ""
			conveyNewRegisterWithInvalidInput(id, emailAddress, givenName, familyName)
		})

		Convey("Given that emailAddress is invalid", func() {
			emailAddress = ""
			conveyNewRegisterWithInvalidInput(id, emailAddress, givenName, familyName)
		})

		Convey("Given that givenName is invalid", func() {
			givenName = ""
			conveyNewRegisterWithInvalidInput(id, emailAddress, givenName, familyName)
		})

		Convey("Given that familyName is invalid", func() {
			familyName = ""
			conveyNewRegisterWithInvalidInput(id, emailAddress, givenName, familyName)
		})
	})
}

func conveyNewRegisterWithInvalidInput(
	id string,
	emailAddress string,
	givenName string,
	familyName string,
) {

	Convey("When a new Register command is created", func() {
		register, err := commands.NewRegister(id, emailAddress, givenName, familyName)

		Convey("It should fail", func() {
			So(err, ShouldBeError)
			So(xerrors.Is(err, shared.ErrInputIsInvalid), ShouldBeTrue)
			So(register, ShouldBeNil)
		})
	})
}

func TestRegisterExposesExpectedValues(t *testing.T) {
	Convey("Given a Register command", t, func() {
		id := "64bcf656-da30-4f5a-b0b5-aead60965aa3"
		emailAddress := "john@doe.com"
		givenName := "John"
		familyName := "Doe"

		idValue, err := values.NewID(id)
		So(err, ShouldBeNil)
		emailAddressValue, err := values.NewEmailAddress(emailAddress)
		So(err, ShouldBeNil)
		personNameValue, err := values.NewPersonName(givenName, familyName)
		So(err, ShouldBeNil)

		register, err := commands.NewRegister(id, emailAddress, givenName, familyName)
		So(err, ShouldBeNil)

		Convey("It should expose the expected values", func() {
			So(idValue.Equals(register.ID()), ShouldBeTrue)
			So(emailAddressValue.Equals(register.EmailAddress()), ShouldBeTrue)
			So(personNameValue.Equals(register.PersonName()), ShouldBeTrue)
			So(register.CommandName(), ShouldEqual, "Register")
			So(idValue.Equals(register.AggregateIdentifier()), ShouldBeTrue)
		})
	})
}