package error

const (
	// ErrorApp represent all application errors
	ErrorApp string = "Application Error"
	// ErrorAppInit initializing app faield
	ErrorAppInit string = "Error App Init"
	// ErrorAppGetAllBooks func allBooks failed
	ErrorAppGetAllBooks string = "Error Get All Books"
	// ErrorAppGetBookByIsbn func bookByISBN failed
	ErrorAppGetBookByIsbn string = "Error Get Book By Isbn"
	// ErrorAppAddBook func addBook failed
	ErrorAppAddBook string = "Error Add New Book"
	// ErrorAppUpdateBook func updateBook failed
	ErrorAppUpdateBook string = "Error Update Book"
	// ErrorAppDeleteBookByIsbn func deleteBook failed
	ErrorAppDeleteBookByIsbn string = "Error Delete Book By Isbn"

	// ErrorDB represent all database errors
	ErrorDB string = "Database Error"
	// ErrorDBSessionInit init DB Session failed
	ErrorDBSessionInit string = "Error Init Database Session"
	// ErrorDBSessionNil DB Session is nil
	ErrorDBSessionNil string = "Error Nil Database Session"
	// ErrorDBIndex DB Indexing failed
	ErrorDBIndex string = "Error Database Indexing"
	// ErrorDBGetAllDocuments func GetAllDocuments failed
	ErrorDBGetAllDocuments string = "Error Database Get All Documents"
	// ErrorDBGetDocumentByKey func GetDocumentByKey failed
	ErrorDBGetDocumentByKey string = "Error Database Get Document By Key"
	// ErrorDBDuplicatedKey func AddDocument failed
	ErrorDBDuplicatedKey string = "Error Database Duplicated Key"
	// ErrorDBUpdateDocument func UpdateDocument failed
	ErrorDBUpdateDocument string = "Error Database Update Document"
	// ErrorDBDeleteDocumentByKey func DeleteDocumentByKey failed
	ErrorDBDeleteDocumentByKey string = "Error Database Delete Document By Key"

	// ErrorSetConfig setting config failed
	ErrorSetConfig string = "Error Set Config"

	// ErrorFileExtension file extension mismatch
	ErrorFileExtension string = "Error File Extension"

	// ErrorJSON represent all JSON errors
	ErrorJSON string = "JSON Error"
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
