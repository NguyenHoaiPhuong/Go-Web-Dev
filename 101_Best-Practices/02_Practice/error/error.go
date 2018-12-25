package error

const (
	// ErrorSetConfig setting config failed
	ErrorSetConfig string = "Error Set Config"

	// ErrorFileExtension file extension mismatch
	ErrorFileExtension string = "Error File Extension"

	// ErrorJSONConvert converting to / from json object failed
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

// Imp implementation
type Imp struct {
	msg string
}

func (e Imp) Error() string {
	return e.msg
}

// SetErrorMessage sets the error message
func (e *Imp) SetErrorMessage(msg string) {
	e.msg = msg
}

// InsertErrorMessage inserts the error message at first
func (e *Imp) InsertErrorMessage(msg string) {
	if len(e.msg) == 0 {
		e.msg = msg
	} else {
		e.msg = msg + "\n" + e.msg
	}
}
