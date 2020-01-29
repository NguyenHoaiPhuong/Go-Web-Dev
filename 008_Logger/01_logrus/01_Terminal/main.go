package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

var logger *log.Logger

func init() {
	logger = log.New()

	// Log as JSON instead of the default ASCII formatter.
	// logger.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	logger.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	logger.SetLevel(log.WarnLevel)

	fmt.Println(log.WarnLevel)
	fmt.Println(log.PanicLevel)
}

func main() {
	logger.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	logger.WithFields(log.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	// log.WithFields(log.Fields{
	// 	"omg":    true,
	// 	"number": 100,
	// }).Fatal("The ice breaks!")

	logger.WithFields(log.Fields{
		"omg":    true,
		"number": 100,
	}).Debug("The ice breaks!")

	// A common pattern is to re-use fields between logging statements by re-using
	// the logrus.Entry returned from WithFields()
	contextLogger := logger.WithFields(log.Fields{
		"common": "this is a common field",
		"other":  "I also should be logged always",
	})

	contextLogger.Info("I'll be logged with common and other field")
	contextLogger.Info("Me too")
}
