

package components

import (
	"errors"
	"fmt"
	"github.com/FastPix/fastpix-go/internal/utils"
	"github.com/FastPix/fastpix-go/optionalnullable"
)

type ErrorDetailsPercentageType string

const (
	ErrorDetailsPercentageTypeInteger ErrorDetailsPercentageType = "integer"
	ErrorDetailsPercentageTypeNumber  ErrorDetailsPercentageType = "number"
)

// ErrorDetailsPercentage - views affected by the specific errors.
type ErrorDetailsPercentage struct {
	Integer *int64   `queryParam:"inline,name=percentage"`
	Number  *float64 `queryParam:"inline,name=percentage"`

	Type ErrorDetailsPercentageType
}

func CreateErrorDetailsPercentageInteger(integer int64) ErrorDetailsPercentage {
	typ := ErrorDetailsPercentageTypeInteger

	return ErrorDetailsPercentage{
		Integer: &integer,
		Type:    typ,
	}
}

func CreateErrorDetailsPercentageNumber(number float64) ErrorDetailsPercentage {
	typ := ErrorDetailsPercentageTypeNumber

	return ErrorDetailsPercentage{
		Number: &number,
		Type:   typ,
	}
}

func (u *ErrorDetailsPercentage) UnmarshalJSON(data []byte) error {

	var integer int64 = int64(0)
	if err := utils.UnmarshalJSON(data, &integer, "", true, nil); err == nil {
		u.Integer = &integer
		u.Type = ErrorDetailsPercentageTypeInteger
		return nil
	}

	var number float64 = float64(0)
	if err := utils.UnmarshalJSON(data, &number, "", true, nil); err == nil {
		u.Number = &number
		u.Type = ErrorDetailsPercentageTypeNumber
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for ErrorDetailsPercentage", string(data))
}

func (u ErrorDetailsPercentage) MarshalJSON() ([]byte, error) {
	if u.Integer != nil {
		return utils.MarshalJSON(u.Integer, "", true)
	}

	if u.Number != nil {
		return utils.MarshalJSON(u.Number, "", true)
	}

	return nil, errors.New("could not marshal union type ErrorDetailsPercentage: all fields are null")
}

type ErrorDetails struct {
	// views affected by the specific errors.
	Percentage optionalnullable.OptionalNullable[ErrorDetailsPercentage] `json:"percentage,omitempty"`
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

func (e *ErrorDetails) GetPercentage() optionalnullable.OptionalNullable[ErrorDetailsPercentage] {
	if e == nil {
		return nil
	}
	return e.Percentage
}

func (e *ErrorDetails) GetNotes() optionalnullable.OptionalNullable[string] {
	if e == nil {
		return nil
	}
	return e.Notes
}

func (e *ErrorDetails) GetMessage() optionalnullable.OptionalNullable[string] {
	if e == nil {
		return nil
	}
	return e.Message
}

func (e *ErrorDetails) GetLastSeen() optionalnullable.OptionalNullable[string] {
	if e == nil {
		return nil
	}
	return e.LastSeen
}

func (e *ErrorDetails) GetID() optionalnullable.OptionalNullable[string] {
	if e == nil {
		return nil
	}
	return e.ID
}

func (e *ErrorDetails) GetDescription() optionalnullable.OptionalNullable[string] {
	if e == nil {
		return nil
	}
	return e.Description
}

func (e *ErrorDetails) GetCount() optionalnullable.OptionalNullable[int64] {
	if e == nil {
		return nil
	}
	return e.Count
}

func (e *ErrorDetails) GetCode() optionalnullable.OptionalNullable[string] {
	if e == nil {
		return nil
	}
	return e.Code
}
