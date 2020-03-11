package customer

import (
	"go-iddd/service/customer/application/writemodel/domain/customer/commands"
	"go-iddd/service/customer/application/writemodel/domain/customer/events"
	"go-iddd/service/customer/application/writemodel/domain/customer/values"
	"go-iddd/service/lib/es"
)

func ChangeEmailAddress(eventStream es.DomainEvents, command commands.ChangeCustomerEmailAddress) es.DomainEvents {
	var emailAddress values.EmailAddress
	var currentStreamVersion uint

	for _, event := range eventStream {
		switch actualEvent := event.(type) {
		case events.CustomerRegistered:
			emailAddress = actualEvent.EmailAddress()
		case events.CustomerEmailAddressChanged:
			emailAddress = actualEvent.EmailAddress()
		}

		currentStreamVersion = event.StreamVersion()
	}

	if emailAddress.Equals(command.EmailAddress()) {
		return nil
	}

	event := events.CustomerEmailAddressWasChanged(
		command.CustomerID(),
		command.EmailAddress(),
		command.ConfirmationHash(),
		currentStreamVersion+1,
	)

	return es.DomainEvents{event}
}