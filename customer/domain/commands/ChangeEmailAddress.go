// Code generated by generate/main.go. DO NOT EDIT.

package commands

import (
	"go-iddd/customer/domain/values"
	"go-iddd/shared"
	"reflect"
	"strings"
)

type ChangeEmailAddress struct {
	id           *values.ID
	emailAddress *values.EmailAddress
}

/*** Factory Method ***/

func NewChangeEmailAddress(
	id string,
	emailAddress string,
) (*ChangeEmailAddress, error) {

	idValue, err := values.RebuildID(id)
	if err != nil {
		return nil, err
	}

	emailAddressValue, err := values.NewEmailAddress(emailAddress)
	if err != nil {
		return nil, err
	}

	changeEmailAddress := &ChangeEmailAddress{
		id:           idValue,
		emailAddress: emailAddressValue,
	}

	return changeEmailAddress, nil
}

/*** Getter Methods ***/

func (changeEmailAddress *ChangeEmailAddress) ID() *values.ID {
	return changeEmailAddress.id
}

func (changeEmailAddress *ChangeEmailAddress) EmailAddress() *values.EmailAddress {
	return changeEmailAddress.emailAddress
}

/*** Implement shared.Command ***/

func (changeEmailAddress *ChangeEmailAddress) AggregateID() shared.IdentifiesAggregates {
	return changeEmailAddress.id
}

func (changeEmailAddress *ChangeEmailAddress) CommandName() string {
	commandType := reflect.TypeOf(changeEmailAddress).String()
	commandTypeParts := strings.Split(commandType, ".")
	commandName := commandTypeParts[len(commandTypeParts)-1]

	return strings.Title(commandName)
}
