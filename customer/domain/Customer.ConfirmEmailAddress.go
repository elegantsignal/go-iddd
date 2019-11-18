package domain

import (
	"go-iddd/customer/domain/commands"
	"go-iddd/customer/domain/events"
	"go-iddd/shared"
)

func ConfirmEmailAddress(customer *Customer, with *commands.ConfirmEmailAddress) shared.DomainEvents {
	if customer.confirmableEmailAddress.IsConfirmed() {
		return nil
	}

	if !customer.confirmableEmailAddress.IsConfirmedBy(with.ConfirmationHash()) {
		event := events.EmailAddressConfirmationHasFailed(
			with.CustomerID(),
			with.ConfirmationHash(),
			customer.currentStreamVersion+1,
		)

		customer.apply(event)

		return shared.DomainEvents{event}
	}

	event := events.EmailAddressWasConfirmed(
		with.CustomerID(),
		with.EmailAddress(),
		customer.currentStreamVersion+1,
	)

	customer.apply(event)

	return shared.DomainEvents{event}
}
