package rabbitmq

import "log"

func logError(message string, err error) {
	if err != nil {
		log.Printf("%s: %s", message, err)
	}
}
