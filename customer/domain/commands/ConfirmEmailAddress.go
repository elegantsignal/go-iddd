// Code generated by generate/main.go. DO NOT EDIT.

package commands

import (
	"go-iddd/customer/domain/values"
	"go-iddd/shared"
	"reflect"
	"strings"
)

type ConfirmEmailAddress struct {
	customerID       *values.CustomerID
	emailAddress     *values.EmailAddress
	confirmationHash *values.ConfirmationHash
}

/*** Factory Method ***/

func NewConfirmEmailAddress(
	customerID string,
	emailAddress string,
	confirmationHash string,
) (*ConfirmEmailAddress, error) {

	customerIDValue, err := values.CustomerIDFrom(customerID)
	if err != nil {
		return nil, err
	}

	emailAddressValue, err := values.EmailAddressFrom(emailAddress)
	if err != nil {
		return nil, err
	}

	confirmationHashValue, err := values.ConfirmationHashFrom(confirmationHash)
	if err != nil {
		return nil, err
	}

	confirmEmailAddress := &ConfirmEmailAddress{
		customerID:       customerIDValue,
		emailAddress:     emailAddressValue,
		confirmationHash: confirmationHashValue,
	}

	return confirmEmailAddress, nil
}

/*** Getter Methods ***/

func (confirmEmailAddress *ConfirmEmailAddress) CustomerID() *values.CustomerID {
	return confirmEmailAddress.customerID
}

func (confirmEmailAddress *ConfirmEmailAddress) EmailAddress() *values.EmailAddress {
	return confirmEmailAddress.emailAddress
}

func (confirmEmailAddress *ConfirmEmailAddress) ConfirmationHash() *values.ConfirmationHash {
	return confirmEmailAddress.confirmationHash
}

/*** Implement shared.Command ***/

func (confirmEmailAddress *ConfirmEmailAddress) AggregateID() shared.IdentifiesAggregates {
	return confirmEmailAddress.customerID
}

func (confirmEmailAddress *ConfirmEmailAddress) CommandName() string {
	commandType := reflect.TypeOf(confirmEmailAddress).String()
	commandTypeParts := strings.Split(commandType, ".")
	commandName := commandTypeParts[len(commandTypeParts)-1]

	return strings.Title(commandName)
}
