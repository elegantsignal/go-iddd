package customer

import (
	"go-iddd/service/lib"
	"go-iddd/service/lib/es"

	"github.com/cockroachdb/errors"
)

func UnmarshalCustomerEvent(name string, payload []byte, streamVersion uint) (es.DomainEvent, error) {
	switch name {
	case "CustomerRegistered":
		return UnmarshalCustomerRegisteredFromJSON(payload, streamVersion), nil
	case "CustomerEmailAddressConfirmed":
		return UnmarshalCustomerEmailAddressConfirmedFromJSON(payload, streamVersion), nil
	case "CustomerEmailAddressChanged":
		return UnmarshalCustomerEmailAddressChangedFromJSON(payload, streamVersion), nil
	}

	err := errors.Mark(
		errors.Wrapf(errors.New("event is unknown"), "unmarshalDomainEvent [%s] failed", name),
		lib.ErrUnmarshalingFailed,
	)

	return nil, err
}