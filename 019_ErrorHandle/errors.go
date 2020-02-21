package errs

import (
	"fmt"
	"strings"
)

// Error struct
type Error struct {
	Code    ErrorCode
	Message string
}

func (err Error) Error() string {
	return fmt.Sprintf("Code: %v, Message: %v", err.Code, err.Message)
}

// Errors : slice of Error
type Errors []Error

// Error get formatted error message
func (errs Errors) Error() string {
	var errors []string
	for _, err := range errs {
		errors = append(errors, err.Error())
	}
	return strings.Join(errors, "; ")
}

// AddError add error to Errors struct
func (errs *Errors) AddError(err Error) {
	if len(err.Message) > 0 {
		*errs = append(*errs, err)
	}
}

// HasError returns true if there is at least one error in the slice
func (errs Errors) HasError() bool {
	if len(errs) > 0 {
		return true
	}
	return false
}

// NewErrors returns a slice of errors
func NewErrors(errors ...Error) Errors {
	errs := make(Errors, 0)
	for _, err := range errors {
		errs.AddError(err)
	}
	return errs
}
