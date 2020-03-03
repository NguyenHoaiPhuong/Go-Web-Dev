package errors

import (
	"log"
)

// CheckError func
func CheckError(err error, msg string) {
	if err != nil {
		log.Fatal(msg+":", err)
	}
}
