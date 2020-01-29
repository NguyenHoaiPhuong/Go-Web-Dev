package main

import (
	"io"
	"io/ioutil"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

// WriterHook is a hook that writes logs of specified LogLevels to specified Writer
type WriterHook struct {
	Writer    io.Writer
	LogLevels []log.Level
}

// Fire will be called when some logging function is called with current hook
// It will format log entry to string and write it to appropriate writer
func (hook *WriterHook) Fire(entry *log.Entry) error {
	line, err := entry.String()
	if err != nil {
		return err
	}
	_, err = hook.Writer.Write([]byte(line))
	return err
}

// Levels define on which log levels this hook would trigger
func (hook *WriterHook) Levels() []log.Level {
	return hook.LogLevels
}

func main() {
	logger := log.New()
	logger.SetOutput(ioutil.Discard)
	logger.SetFormatter(&log.TextFormatter{
		ForceColors: true,
	})
	logger.SetLevel(log.DebugLevel)
	writerHook := &WriterHook{ // Send logs with level higher than warning to stderr
		Writer: os.Stdout,
		LogLevels: []log.Level{
			log.PanicLevel,
			log.FatalLevel,
			log.ErrorLevel,
			log.WarnLevel,
			log.InfoLevel,
			// log.DebugLevel,
		},
	}
	logger.AddHook(writerHook)

	logger.WithField("function", "main").Warn("Start ...")

	for i := 0; i < 3; i++ {
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
