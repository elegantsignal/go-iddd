package customer

import (
	"github.com/AntonStoeckl/go-iddd/service/customeraccounts/application/domain"
	"github.com/AntonStoeckl/go-iddd/service/shared/es"
)

func Delete(eventStream es.EventStream) es.RecordedEvents {
	customer := buildCurrentStateFrom(eventStream)

	if err := assertNotDeleted(customer); err != nil {
		return nil
	}

	event := domain.BuildCustomerDeleted(
		customer.id,
		customer.emailAddress,
		customer.currentStreamVersion+1,
	)

	return es.RecordedEvents{event}
}