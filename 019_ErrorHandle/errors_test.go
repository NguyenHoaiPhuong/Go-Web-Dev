package errs

import (
	"fmt"
	"testing"
)

func backendCreateUser() {
	errs := entrystoreCreateUser()
	if errs.HasError() {
		fmt.Println(errs.Error())
	}
}

func entrystoreCreateUser() Errors {
	errs := mgoCreateUser()
	if errs.HasError() {
		errs.AddError(Error{Code: ErrCodeAddressExisted, Message: ErrMsgInvalidName})
		return errs
	}
	return nil
}

func mgoCreateUser() Errors {
	return NewErrors(Error{Code: ErrCodeMongoCreate, Message: "cannot create model"})
}

func TestError(t *testing.T) {
	backendCreateUser()
}

func TestErrorCode(t *testing.T) {
	fmt.Printf("Error code: %d, Error message: %s", ErrCodeInvalidEmail, ErrMsgInvalidEmail)
	fmt.Printf("Error code: %d, Error message: %s", ErrCodeInvalidPassword, ErrMsgInvalidPassword)
}
