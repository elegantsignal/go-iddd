package events_test

import (
	"go-iddd/customer/domain/events"
	"go-iddd/customer/domain/values"
	"go-iddd/shared"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
	"golang.org/x/xerrors"
)

func TestEmailAddressWasChanged(t *testing.T) {
	Convey("Given valid parameters as input", t, func() {
		id := values.GenerateID()
		emailAddress, err := values.NewEmailAddress("foo@bar.com")
		So(err, ShouldBeNil)

		Convey("When a new EmailAddressChanged event is created", func() {
			streamVersion := uint(1)
			emailAddressChanged := events.EmailAddressWasChanged(id, emailAddress, streamVersion)

			Convey("It should succeed", func() {
				So(emailAddressChanged, ShouldNotBeNil)
				So(emailAddressChanged, ShouldImplement, (*shared.DomainEvent)(nil))
				So(emailAddressChanged, ShouldHaveSameTypeAs, (*events.EmailAddressChanged)(nil))
			})
		})
	})
}

func TestEmailAddressChangedExposesExpectedValues(t *testing.T) {
	Convey("Given a EmailAddressChanged event", t, func() {
		id := values.GenerateID()
		emailAddress, err := values.NewEmailAddress("foo@bar.com")
		So(err, ShouldBeNil)
		streamVersion := uint(1)

		beforeItOccurred := time.Now()
		emailAddressChanged := events.EmailAddressWasChanged(id, emailAddress, streamVersion)
		afterItOccurred := time.Now()

		Convey("It should expose the expected values", func() {
			So(emailAddressChanged.ID(), ShouldResemble, id)
			So(emailAddressChanged.EmailAddress(), ShouldResemble, emailAddress)
			So(emailAddressChanged.Identifier(), ShouldEqual, id.String())
			So(emailAddressChanged.EventName(), ShouldEqual, "CustomerEmailAddressChanged")
			itOccurred, err := time.Parse(shared.DomainEventMetaTimestampFormat, emailAddressChanged.OccurredAt())
			So(err, ShouldBeNil)
			So(beforeItOccurred, ShouldHappenBefore, itOccurred)
			So(afterItOccurred, ShouldHappenAfter, itOccurred)
			So(emailAddressChanged.StreamVersion(), ShouldEqual, streamVersion)
		})
	})
}

func TestEmailAddressChangedMarshalJSON(t *testing.T) {
	Convey("Given a EmailAddressChanged event", t, func() {
		id := values.GenerateID()
		emailAddress, err := values.NewEmailAddress("foo@bar.com")
		So(err, ShouldBeNil)
		streamVersion := uint(1)

		emailAddressChanged := events.EmailAddressWasChanged(id, emailAddress, streamVersion)

		Convey("When it is marshaled to json", func() {
			data, err := emailAddressChanged.MarshalJSON()

			Convey("It should create the expected json", func() {
				So(err, ShouldBeNil)
				So(string(data), ShouldStartWith, "{")
				So(string(data), ShouldEndWith, "}")
			})
		})
	})
}

func TestEmailAddressChangedUnmarshalJSON(t *testing.T) {
	Convey("Given a EmailAddressChanged event marshaled to json", t, func() {
		id := values.GenerateID()
		emailAddress, err := values.NewEmailAddress("foo@bar.com")
		So(err, ShouldBeNil)
		streamVersion := uint(1)

		emailAddressChanged := events.EmailAddressWasChanged(id, emailAddress, streamVersion)

		data, err := emailAddressChanged.MarshalJSON()

		Convey("And when it is unmarshaled", func() {
			unmarshaled := &events.EmailAddressChanged{}
			err := unmarshaled.UnmarshalJSON(data)

			Convey("It should be equal to the original EmailAddressChanged event", func() {
				So(err, ShouldBeNil)
				So(emailAddressChanged, ShouldResemble, unmarshaled)
			})
		})
	})

	Convey("Given invalid json", t, func() {
		data := []byte("666")

		Convey("When it is unmarshaled to EmailAddressChanged event", func() {
			unmarshaled := &events.EmailAddressChanged{}
			err := unmarshaled.UnmarshalJSON(data)

			Convey("It should fail", func() {
				So(err, ShouldNotBeNil)
				So(xerrors.Is(err, shared.ErrUnmarshalingFailed), ShouldBeTrue)
			})
		})
	})
}
