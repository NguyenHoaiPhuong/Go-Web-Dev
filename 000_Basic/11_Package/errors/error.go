package errors

// Error includes Code and Message
type Error struct {
	code    int32
	message string
	service string
}

// Code : error code
func (e *Error) Code() int32 {
	return e.code
}

// Message : error message
func (e *Error) Message() string {
	return e.message
}

// Service : service name
func (e *Error) Service() string {
	return e.service
}

// New : return a new pointer of Error
func New(srvName string, code int32, message string) *Error {
	e := new(Error)
	e.code = code
	e.message = message
	e.service = srvName

	return e
}
