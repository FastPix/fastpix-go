

package components

import (
	"errors"
	"fmt"
	"github.com/FastPix/fastpix-go/internal/utils"
	"github.com/FastPix/fastpix-go/optionalnullable"
)

type TopErrorDetailsPercentageType string

const (
	TopErrorDetailsPercentageTypeInteger TopErrorDetailsPercentageType = "integer"
	TopErrorDetailsPercentageTypeNumber  TopErrorDetailsPercentageType = "number"
)

// TopErrorDetailsPercentage - views affected by the specific errors.
type TopErrorDetailsPercentage struct {
	Integer *int64   `queryParam:"inline,name=percentage"`
	Number  *float64 `queryParam:"inline,name=percentage"`

	Type TopErrorDetailsPercentageType
}

func CreateTopErrorDetailsPercentageInteger(integer int64) TopErrorDetailsPercentage {
	typ := TopErrorDetailsPercentageTypeInteger

	return TopErrorDetailsPercentage{
		Integer: &integer,
		Type:    typ,
	}
}

func CreateTopErrorDetailsPercentageNumber(number float64) TopErrorDetailsPercentage {
	typ := TopErrorDetailsPercentageTypeNumber

	return TopErrorDetailsPercentage{
		Number: &number,
		Type:   typ,
	}
}

func (u *TopErrorDetailsPercentage) UnmarshalJSON(data []byte) error {

	var integer int64 = int64(0)
	if err := utils.UnmarshalJSON(data, &integer, "", true, nil); err == nil {
		u.Integer = &integer
		u.Type = TopErrorDetailsPercentageTypeInteger
		return nil
	}

	var number float64 = float64(0)
	if err := utils.UnmarshalJSON(data, &number, "", true, nil); err == nil {
		u.Number = &number
		u.Type = TopErrorDetailsPercentageTypeNumber
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for TopErrorDetailsPercentage", string(data))
}

func (u TopErrorDetailsPercentage) MarshalJSON() ([]byte, error) {
	if u.Integer != nil {
		return utils.MarshalJSON(u.Integer, "", true)
	}

	if u.Number != nil {
		return utils.MarshalJSON(u.Number, "", true)
	}

	return nil, errors.New("could not marshal union type TopErrorDetailsPercentage: all fields are null")
}

type UniqueViewersEffectedPercentageType string

const (
	UniqueViewersEffectedPercentageTypeInteger UniqueViewersEffectedPercentageType = "integer"
	UniqueViewersEffectedPercentageTypeNumber  UniqueViewersEffectedPercentageType = "number"
)

// UniqueViewersEffectedPercentage - percentage of unique viewers affected by the specific error.
type UniqueViewersEffectedPercentage struct {
	Integer *int64   `queryParam:"inline,name=uniqueViewersEffectedPercentage"`
	Number  *float64 `queryParam:"inline,name=uniqueViewersEffectedPercentage"`

	Type UniqueViewersEffectedPercentageType
}

func CreateUniqueViewersEffectedPercentageInteger(integer int64) UniqueViewersEffectedPercentage {
	typ := UniqueViewersEffectedPercentageTypeInteger

	return UniqueViewersEffectedPercentage{
		Integer: &integer,
		Type:    typ,
	}
}

func CreateUniqueViewersEffectedPercentageNumber(number float64) UniqueViewersEffectedPercentage {
	typ := UniqueViewersEffectedPercentageTypeNumber

	return UniqueViewersEffectedPercentage{
		Number: &number,
		Type:   typ,
	}
}

func (u *UniqueViewersEffectedPercentage) UnmarshalJSON(data []byte) error {

	var integer int64 = int64(0)
	if err := utils.UnmarshalJSON(data, &integer, "", true, nil); err == nil {
		u.Integer = &integer
		u.Type = UniqueViewersEffectedPercentageTypeInteger
		return nil
	}

	var number float64 = float64(0)
	if err := utils.UnmarshalJSON(data, &number, "", true, nil); err == nil {
		u.Number = &number
		u.Type = UniqueViewersEffectedPercentageTypeNumber
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for UniqueViewersEffectedPercentage", string(data))
}

func (u UniqueViewersEffectedPercentage) MarshalJSON() ([]byte, error) {
	if u.Integer != nil {
		return utils.MarshalJSON(u.Integer, "", true)
	}

	if u.Number != nil {
		return utils.MarshalJSON(u.Number, "", true)
	}

	return nil, errors.New("could not marshal union type UniqueViewersEffectedPercentage: all fields are null")
}

type TopErrorDetails struct {
	// views affected by the specific errors.
	Percentage optionalnullable.OptionalNullable[TopErrorDetailsPercentage] `json:"percentage,omitempty"`
	// percentage of unique viewers affected by the specific error.
	UniqueViewersEffectedPercentage optionalnullable.OptionalNullable[UniqueViewersEffectedPercentage] `json:"uniqueViewersEffectedPercentage,omitempty"`
	// Information about the specific error.
	Notes optionalnullable.OptionalNullable[string] `json:"notes,omitempty"`
	// error message or description.
	Message optionalnullable.OptionalNullable[string] `json:"message,omitempty"`
	// The timestamp of when the error was last observed.
	LastSeen optionalnullable.OptionalNullable[string] `json:"lastSeen,omitempty"`
	// unique identifier for the specific error.
	ID optionalnullable.OptionalNullable[string] `json:"id,omitempty"`
	// description of the specific error.
	Description optionalnullable.OptionalNullable[string] `json:"description,omitempty"`
	// Number of occurrences of the specific error.
	Count optionalnullable.OptionalNullable[int64] `json:"count,omitempty"`
	// Error code associated with the specific error.
	Code optionalnullable.OptionalNullable[string] `json:"code,omitempty"`
}

func (t *TopErrorDetails) GetPercentage() optionalnullable.OptionalNullable[TopErrorDetailsPercentage] {
	if t == nil {
		return nil
	}
	return t.Percentage
}

func (t *TopErrorDetails) GetUniqueViewersEffectedPercentage() optionalnullable.OptionalNullable[UniqueViewersEffectedPercentage] {
	if t == nil {
		return nil
	}
	return t.UniqueViewersEffectedPercentage
}

func (t *TopErrorDetails) GetNotes() optionalnullable.OptionalNullable[string] {
	if t == nil {
		return nil
	}
	return t.Notes
}

func (t *TopErrorDetails) GetMessage() optionalnullable.OptionalNullable[string] {
	if t == nil {
		return nil
	}
	return t.Message
}

func (t *TopErrorDetails) GetLastSeen() optionalnullable.OptionalNullable[string] {
	if t == nil {
		return nil
	}
	return t.LastSeen
}

func (t *TopErrorDetails) GetID() optionalnullable.OptionalNullable[string] {
	if t == nil {
		return nil
	}
	return t.ID
}

func (t *TopErrorDetails) GetDescription() optionalnullable.OptionalNullable[string] {
	if t == nil {
		return nil
	}
	return t.Description
}

func (t *TopErrorDetails) GetCount() optionalnullable.OptionalNullable[int64] {
	if t == nil {
		return nil
	}
	return t.Count
}

func (t *TopErrorDetails) GetCode() optionalnullable.OptionalNullable[string] {
	if t == nil {
		return nil
	}
	return t.Code
}
