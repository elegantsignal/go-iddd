package valueobjects

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"go-iddd/shared"
	"math/rand"
	"strconv"
	"time"

	"golang.org/x/xerrors"
)

type ConfirmationHash struct {
	value string
}

/*** Factory methods ***/

func GenerateConfirmationHash(using string) *ConfirmationHash {
	randomInt := rand.New(rand.NewSource(time.Now().UnixNano())).Int()
	md5Sum := md5.Sum([]byte(strconv.Itoa(randomInt) + using))
	value := fmt.Sprintf("%x", md5Sum)

	return buildConfirmationHash(value)
}

func ReconstituteConfirmationHash(from string) *ConfirmationHash {
	return buildConfirmationHash(from)
}

func buildConfirmationHash(from string) *ConfirmationHash {
	return &ConfirmationHash{value: from}
}

/*** Public methods implementing ConfirmationHash ***/

func (confirmationHash *ConfirmationHash) Hash() string {
	return confirmationHash.value
}

func (confirmationHash *ConfirmationHash) MustMatch(other *ConfirmationHash) error {
	if confirmationHash.Hash() != other.Hash() {
		return xerrors.Errorf("confirmationHash.MustMatch: input does not match: %w", shared.ErrInvalidInput) // TODO: use a distinct error type?
	}

	return nil
}

/*** Implement json.Marshaler ***/

func (confirmationHash *ConfirmationHash) MarshalJSON() ([]byte, error) {
	bytes, err := json.Marshal(confirmationHash.value)
	if err != nil {
		return bytes, xerrors.Errorf("confirmationHash.MarshalJSON: %s: %w", err, shared.ErrMarshaling)
	}

	return bytes, nil
}

/*** Implement json.Unmarshaler ***/

func (confirmationHash *ConfirmationHash) UnmarshalJSON(data []byte) error {
	var value string

	if err := json.Unmarshal(data, &value); err != nil {
		return xerrors.Errorf("confirmationHash.UnmarshalJSON: %s: %w", err, shared.ErrUnmarshaling)
	}

	confirmationHash.value = value

	return nil
}
