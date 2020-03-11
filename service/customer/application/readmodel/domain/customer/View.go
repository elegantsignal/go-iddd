package customer

import (
	"go-iddd/service/customer/application/readmodel/domain/customer/events"
	"go-iddd/service/lib/es"
)

type View struct {
	ID                      string
	EmailAddress            string
	IsEmailAddressConfirmed bool
	GivenName               string
	FamilyName              string
	Version                 uint
}

func BuildViewFrom(eventStream es.DomainEvents) View {
	customerView := View{}

	for _, event := range eventStream {
		switch actualEvent := event.(type) {
		case events.CustomerRegistered:
			customerView.ID = actualEvent.CustomerID()
			customerView.EmailAddress = actualEvent.EmailAddress()
			customerView.GivenName = actualEvent.GivenName()
			customerView.FamilyName = actualEvent.FamilyName()
		case events.CustomerEmailAddressConfirmed:
			customerView.IsEmailAddressConfirmed = true
		case events.CustomerEmailAddressChanged:
			customerView.EmailAddress = actualEvent.EmailAddress()
			customerView.IsEmailAddressConfirmed = false
		}

		customerView.Version = event.StreamVersion()
	}

	return customerView
}
