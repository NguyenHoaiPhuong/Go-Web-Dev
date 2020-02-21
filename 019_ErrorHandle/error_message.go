package errs

/************ Backend error messages ************/
const (
	// ErrMsgInvalidEmail : error message invalid email
	ErrMsgInvalidEmail string = "invalid email"

	// ErrMsgInvalidPassword : error message invalid password
	ErrMsgInvalidPassword string = "invalid password"

	// ErrMsgEmailNotExists : error message email doesn't exist
	ErrMsgEmailNotExists string = "email doesn't exist"

	// ErrMsgEmailAlreadyExists : error message email already exists
	ErrMsgEmailAlreadyExists string = "email already exists"

	// ErrMsgPasswordMismatch : perror message assword and confirm password must match
	ErrMsgPasswordMismatch string = "password and confirm password must match"

	// ErrMsgInvalidCredentials : error message invalid credentials
	ErrMsgInvalidCredentials string = "invalid credentials"

	// ErrMsgInvalidName : error message invalid name
	ErrMsgInvalidName string = "invalid name"

	// ErrMsgUpdateUser : error message cannot update user profile
	ErrMsgUpdateUser string = "cannot update user profile"
)

/************ Entry-cache error messages ************/
const ()

/************ Entry-store error messages ************/
const (
	// ErrMsgAddressExisted : error message user address is existed
	ErrMsgAddressExisted string = "user address is existed"

	// ErrMsgAddressSaveFail : user address save fail
	ErrMsgAddressSaveFail string = "user address save fail"
)
