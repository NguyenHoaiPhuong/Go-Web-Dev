package main

import (
	"os"
	"time"

	"github.com/NguyenHoaiPhuong/Go-Web-Dev/008_Logger/04_Mongo/config"

	log "github.com/sirupsen/logrus"
)

func main() {
	cfg := config.SetupAndGetConfig("./resources/conf.json")

	// stdLogger := log.New()
	// stdLogger.SetOutput(os.Stdout)
	// stdLogger.SetLevel(log.AllLevels[*cfg.LoggingLevel])

	writerHook := &WriterHook{ // Send logs with level higher than warning to stderr
		Writer: os.Stdout,
		LogLevels: []log.Level{
			log.PanicLevel,
			log.FatalLevel,
			log.ErrorLevel,
			log.WarnLevel,
			log.InfoLevel,
			log.DebugLevel,
		},
	}
	logger.AddHook(writerHook)

	mongoHook := new(mongodbHook)
	mongoHook.Init(cfg.MongoDBConfig)
	logger.AddHook(mongoHook)

	logger.WithField("function", "main").Warn("Start ...")

	for i := 0; i < 10; i++ {
		logger.WithFields(log.Fields{
			"animal": "walrus",
			"number": 1,
			"size":   10,
		}).Info("A walrus appears")
		time.Sleep(time.Second)
	}

	logger.WithField("function", "main").Debug("Debug stop ...")
	logger.WithField("function", "main").Trace("Trace stop ...")
}
