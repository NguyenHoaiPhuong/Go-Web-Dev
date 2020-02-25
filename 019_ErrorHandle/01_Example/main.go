package main

import (
	"fmt"
)

var (
	// ErrorMongoCreate : cannot create modelMessage
	ErrorMongoCreate *Error = New(1, "cannot create model")

	// ErrorCreateUser : cannot create user
	ErrorCreateUser *Error = New(2, "cannot create user")

	// ErrorInvalidUsername : invalid username
	ErrorInvalidUsername *Error = New(3, "invalid username")

	// ErrorInvalidPassword : invalid password
	ErrorInvalidPassword *Error = New(4, "invalid password")

	// ErrorPasswordMismatched : password and confirmed password mismatched
	ErrorPasswordMismatched *Error = New(5, "password and confirmed password mismatched")
)

func testErrors() {
	fmt.Println("Test errors receiver functions")

	err := ErrorMongoCreate
	fmt.Println("err:", err.ErrorWithCode())

	err.WithMessage("update")
	fmt.Println("err:", err.ErrorWithCode())

	err.WithMessage("cannot create user")
	fmt.Println("err:", err.ErrorWithCode())
}

func main() {
	testErrors()
	// testCompareError()
}
