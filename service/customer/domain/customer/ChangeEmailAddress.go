package customer

import (
	"github.com/AntonStoeckl/go-iddd/service/customer/domain/customer/commands"
	"github.com/AntonStoeckl/go-iddd/service/customer/domain/customer/events"
	"github.com/AntonStoeckl/go-iddd/service/lib/es"
	"github.com/cockroachdb/errors"
)

func ChangeEmailAddress(eventStream es.DomainEvents, command commands.ChangeCustomerEmailAddress) (es.DomainEvents, error) {
	state := buildCustomerStateFrom(eventStream)

	if err := MustNotBeDeleted(state); err != nil {
		return nil, errors.Wrap(err, "changeEmailAddress")
	}

	if state.emailAddress.Equals(command.EmailAddress()) {
		return nil, nil
	}

	event := events.CustomerEmailAddressWasChanged(
		state.id,
		command.EmailAddress(),
		command.ConfirmationHash(),
		state.emailAddress,
		state.currentStreamVersion+1,
	)

	return es.DomainEvents{event}, nil
}
