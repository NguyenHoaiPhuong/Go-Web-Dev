package errs

type (
	// ErrorCode : error code
	ErrorCode uint
)

/************ Third party error code from 000 to 999 ************/
const (
	/************ Mongodb error ************/
	// ErrCodeMongoConnection : mongodb connection error
	ErrCodeMongoConnection ErrorCode = iota

	// ErrCodeMongoCreate : mongodb create model error
	ErrCodeMongoCreate

	// ErrCodeMongoRead : mongodb read model error
	ErrCodeMongoRead

	// ErrCodeMongoUpdate : mongodb update model error
	ErrCodeMongoUpdate

	// ErrCodeMongoDelete : mongodb delete model error
	ErrCodeMongoDelete

	/************ Redis error ************/

	/************ Rabbitmq error ************/
)

/************ Backend error code from 1000 to 1999 ************/
const (
	// ErrCodeInvalidEmail : invalid email
	ErrCodeInvalidEmail ErrorCode = iota + 1000

	// ErrCodeInvalidPassword : invalid password
	ErrCodeInvalidPassword

	// ErrCodeEmailNotExists : email doesn't exist
	ErrCodeEmailNotExists

	// ErrCodeEmailAlreadyExists : email already exists
	ErrCodeEmailAlreadyExists

	// ErrCodePasswordMismatch : password and confirm password must match
	ErrCodePasswordMismatch

	// ErrCodeInvalidCredentials : invalid credentials
	ErrCodeInvalidCredentials

	// ErrCodeInvalidName : invalid name
	ErrCodeInvalidName

	// ErrCodeUpdateUser : cannot update user profile
	ErrCodeUpdateUser
)

/************ Entry-cache error code from 2000 to 29999 ************/
const ()

/************ Entry-store error code from 3000 to 39999 ************/
const (
	// ErrCodeAddressExisted : error message user address is existed
	ErrCodeAddressExisted ErrorCode = iota + 3000

	// ErrCodeAddressSaveFail : user address save fail
	ErrCodeAddressSaveFail
)

/************ Assets service error code from 4000 to 49999 ************/
const ()

/************ Email service error code from 5000 to 59999 ************/
const ()

/************ Scan service error code from 6000 to 69999 ************/
const ()

/************ Event service error code from 7000 to 79999 ************/
const ()

/************ Push notification error code from 8000 to 89999 ************/
const ()

/************ bo controller error code from 9000 to 99999 ************/
const ()
