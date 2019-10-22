package main

import (
	"database/sql"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

type dbHook struct {
	*sql.DB
}

func (dbHook) Levels() []logrus.Level {
	return []log.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
		logrus.InfoLevel,
		logrus.DebugLevel,
	}
}

func (db dbHook) Fire(e *log.Entry) error {
	_, err := db.Exec("insert into log (time, level, message) values (?, ?, ?)",
		e.Time,
		e.Level,
		e.Message,
	)

	return err
}

var logger *log.Logger

func init() {
	logger = log.New()
	logger.SetOutput(os.Stdout)
	logger.SetLevel(log.ErrorLevel)
}

func main() {
	entry := logger.WithField("function", "main")
	entry.Info("Start ...")

	db, err := sql.Open("sqlite3", "log.db")
	if err != nil {
		entry.Fatalf("Connecting sqlite3 error: %v", err)
	}

	logger.AddHook(dbHook{db})

	_, err = db.Exec("create table log (time, level, message)")
	if err != nil {
		entry.Errorf("Creating table log  error: %v", err)
	}

	for i := 0; i < 10; i++ {
		logger.WithFields(logrus.Fields{
			"animal": "walrus",
			"number": 1,
			"size":   10,
		}).Info("A walrus appears")
		time.Sleep(time.Second)
	}
}
