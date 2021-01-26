package cashbookerror

import (
	"errors"
	"fmt"
)

// APIError is an error object for the frontend
type APIError struct {
	Code int64
	message string
}

// NewAPIError creates a new APIError
func NewAPIError(code int64, message string) *APIError {
	return &APIError{
		Code: code,
		message: message,
	}
}

// AccountCodeError type represents an related to Account Code type
type AccountCodeError struct {
	Code int64
	Err error
}

func (e *AccountCodeError) Error() string {
	return e.Err.Error()
}

// NewAccountCodeError creates a AccountCode error with all params
func NewAccountCodeError(code int64, message string) error {
	return &AccountCodeError{
		Code: code,
		Err: errors.New(message),
	}
}

// AccountCodeDoesNotExistError creates an account code error with code 404
func AccountCodeDoesNotExistError(code string) error {
	return &AccountCodeError{
		Code: 404,
		Err: fmt.Errorf("Account code with code %s does not exist", code),
	}
}

// UnknownAccountCodeError creates an account code error with code 500
// Also it stores the actual error along with it
func UnknownAccountCodeError(e error) error {
	return &AccountCodeError{
		Code: 500,
		Err: e,
	}
}

// DoesAccountCodeExist tells whether the error tells about non existing account code
func (e *AccountCodeError) DoesAccountCodeExist() bool {
	return e.Code == 404
}

