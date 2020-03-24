package main

import (
	"os"

	"time"

	"github.com/robfig/cron"
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
	logger.SetLevel(log.DebugLevel)
}

func main() {
	logger.Info("Start")

	go func() {
		tnxCronInstance := cron.NewWithLocation(time.UTC)
		// run at 7:44 every Tuesday, GMT +0
		tnxCronInstance.AddFunc("44 7 * * *", func() {
			logger.Info("run at 7:44 every Tuesday")
		})
		tnxCronInstance.Start()
	}()

	select {}
}
