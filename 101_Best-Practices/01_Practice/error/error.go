package error

const (
	// ErrorApplicationInit Application Initialization failed
	ErrorApplicationInit string = "Error Application Initialization"

	// ErrorDatabaseConnection DB connection failed
	ErrorDatabaseConnection string = "Error Database Connection"

	// ErrorSetConfig setting config failed
	ErrorSetConfig string = "Error Set Config"

	// ErrorFileExtension file extension mismatch
	ErrorFileExtension string = "Error File Extension"

	// ErrorJSONConvert converting to json byte slice failed
	ErrorJSONConvert string = "Error Json Convert"
	// ErrorJSONRead reading json file failed
	ErrorJSONRead string = "Error Json Read"
	// ErrorJSONWrite writing to json file failed
	ErrorJSONWrite string = "Error Json Write"
)

// Error interface
type Error interface {
	Error() string
}

// ErrorImp implementation
type ErrorImp struct {
	msg string
}

func (e ErrorImp) Error() string {
	return e.msg
}

// SetErrorMessage sets the error message
func (e *ErrorImp) SetErrorMessage(msg string) {
	e.msg = msg
}

// InsertErrorMessage inserts the error message at first
func (e *ErrorImp) InsertErrorMessage(msg string) {
	if len(e.msg) == 0 {
		e.msg = msg
	} else {
		e.msg = msg + "\n" + e.msg
	}
}
