package main

import (
	"fmt"

	"github.com/pkg/errors"
)

// Error struct
type Error struct {
	code uint
	err  error
}

// New create a new pointer of Error
func New(code uint, message string) *Error {
	e := new(Error)
	e.code = code
	e.err = errors.New(message)
	return e
}

// GetCode returns error code
func (e *Error) GetCode() uint {
	return e.code
}

// SetCode :set error code
func (e *Error) SetCode(code uint) {
	e.code = code
}

func (e *Error) Error() string {
	return e.err.Error()
}

// func (err Error) Error() string {
// 	return fmt.Sprintf("Code: %v, Message: %v", err.Code, err.Message)
// }

// ErrorWithCode : print error code and message
func (e *Error) ErrorWithCode() string {
	return fmt.Sprintf("Code: %v, Message: %v", e.code, e.Error())
}

// WithMessage : annotates error with a new message
func (e *Error) WithMessage(msg string) {
	e.err = errors.WithMessage(e.err, msg)
}

// Wrap : annotates error with a new message
func (e *Error) Wrap(msg string) {
	e.err = errors.Wrap(e.err, msg)
}
