package errs

import (
	"fmt"
	"testing"
)

func TestError(t *testing.T) {
	ErrInvalidEmail := Error{Code: 1000, Message: "invalid email"}
	ErrInvalidPassword := Error{Code: 1001, Message: "invalid password"}
	ErrEmailNotExists := Error{Code: 1002, Message: "email doesn't exist"}
	ErrEmailAlreadyExists := Error{Code: 1003, Message: "email already exists"}

	errs := NewErrors(ErrInvalidEmail, ErrInvalidPassword, ErrEmailNotExists, ErrEmailAlreadyExists)
	fmt.Println(errs.Error())
}
