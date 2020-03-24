package main

import (
	"fmt"
	"os"

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
		tnxCronInstance := cron.New()
		tnxCronInstance.AddFunc(fmt.Sprintf("@every %s", "5s"), func() {
			logger.Info(fmt.Sprintf("scan @every %s", "5s"))
		})
		tnxCronInstance.Start()
	}()

	select {}

}
