package domain

import (
	"github.com/AntonStoeckl/go-iddd/src/customeraccounts/hexagon/application/domain/customer/value"
	"github.com/AntonStoeckl/go-iddd/src/shared/es"
)

type CustomerDeleted struct {
	customerID   value.CustomerID
	emailAddress value.EmailAddress
	meta         es.EventMeta
}

func BuildCustomerDeleted(
	customerID value.CustomerID,
	emailAddress value.EmailAddress,
	causationID es.MessageID,
	streamVersion uint,
) CustomerDeleted {

	event := CustomerDeleted{
		customerID:   customerID,
		emailAddress: emailAddress,
	}

	event.meta = es.BuildEventMeta(event, causationID, streamVersion)

	return event
}

func RebuildCustomerDeleted(
	customerID string,
	emailAddress string,
	meta es.EventMeta,
) CustomerDeleted {

	event := CustomerDeleted{
		customerID:   value.RebuildCustomerID(customerID),
		emailAddress: value.RebuildEmailAddress(emailAddress),
		meta:         meta,
	}

	return event
}

func (event CustomerDeleted) CustomerID() value.CustomerID {
	return event.customerID
}

func (event CustomerDeleted) EmailAddress() value.EmailAddress {
	return event.emailAddress
}

func (event CustomerDeleted) Meta() es.EventMeta {
	return event.meta
}

func (event CustomerDeleted) IsFailureEvent() bool {
	return false
}

func (event CustomerDeleted) FailureReason() error {
	return nil
}
