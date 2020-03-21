package events

import (
	"github.com/AntonStoeckl/go-iddd/service/customer/domain/customer/values"
	jsoniter "github.com/json-iterator/go"
)

type CustomerEmailAddressConfirmed struct {
	customerID   values.CustomerID
	emailAddress values.EmailAddress
	meta         EventMeta
}

func CustomerEmailAddressWasConfirmed(
	customerID values.CustomerID,
	emailAddress values.EmailAddress,
	streamVersion uint,
) CustomerEmailAddressConfirmed {

	event := CustomerEmailAddressConfirmed{
		customerID:   customerID,
		emailAddress: emailAddress,
	}

	event.meta = BuildEventMeta(event, streamVersion)

	return event
}

func (event CustomerEmailAddressConfirmed) CustomerID() values.CustomerID {
	return event.customerID
}

func (event CustomerEmailAddressConfirmed) EmailAddress() values.EmailAddress {
	return event.emailAddress
}

func (event CustomerEmailAddressConfirmed) EventName() string {
	return event.meta.eventName
}

func (event CustomerEmailAddressConfirmed) OccurredAt() string {
	return event.meta.occurredAt
}

func (event CustomerEmailAddressConfirmed) StreamVersion() uint {
	return event.meta.streamVersion
}

func (event CustomerEmailAddressConfirmed) MarshalJSON() ([]byte, error) {
	data := &struct {
		CustomerID   string    `json:"customerID"`
		EmailAddress string    `json:"emailAddress"`
		Meta         EventMeta `json:"meta"`
	}{
		CustomerID:   event.customerID.ID(),
		EmailAddress: event.emailAddress.EmailAddress(),
		Meta:         event.meta,
	}

	return jsoniter.ConfigFastest.Marshal(data)
}

func UnmarshalCustomerEmailAddressConfirmedFromJSON(
	data []byte,
	streamVersion uint,
) CustomerEmailAddressConfirmed {

	anyData := jsoniter.ConfigFastest.Get(data)

	event := CustomerEmailAddressConfirmed{
		customerID:   values.RebuildCustomerID(anyData.Get("customerID").ToString()),
		emailAddress: values.RebuildEmailAddress(anyData.Get("emailAddress").ToString()),
		meta:         UnmarshalEventMetaFromJSON(data, streamVersion),
	}

	return event
}
