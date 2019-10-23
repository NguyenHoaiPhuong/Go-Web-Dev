package main

import (
	"os"
	"time"

	"github.com/NguyenHoaiPhuong/Go-Web-Dev/008_Logger/04_Mongo/config"
	log "github.com/sirupsen/logrus"
)

var logger *log.Logger

func init() {
	logger = log.New()
	logger.SetOutput(os.Stdout)
	logger.SetLevel(log.ErrorLevel)
}

func main() {
	cfg := config.SetupAndGetConfig("./resources/conf.json")

	entry := logger.WithField("function", "main")
	entry.Info("Start ...")

	hook := new(mongodbHook)
	hook.Init(cfg.MongodbConfig)
	logger.AddHook(hook)

	for i := 0; i < 10; i++ {
		logger.WithFields(log.Fields{
			"animal": "walrus",
			"number": 1,
			"size":   10,
		}).Info("A walrus appears")
		time.Sleep(time.Second)
	}
}
