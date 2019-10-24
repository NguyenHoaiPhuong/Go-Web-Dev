package main

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/NguyenHoaiPhuong/Go-Web-Dev/008_Logger/04_Mongo/config"
	"github.com/globalsign/mgo"
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

type mongodbHook struct {
	cfg        *config.MongoDBConfig
	session    *mgo.Session
	database   *mgo.Database
	collection *mgo.Collection

	LogLevels []log.Level
}

func (hook *mongodbHook) Init(cfg *config.MongoDBConfig) {
	hook.cfg = cfg.Copy()
	hook.initDatabase()
	hook.LogLevels = []log.Level{
		log.PanicLevel,
		log.FatalLevel,
		log.ErrorLevel,
		log.WarnLevel,
		log.InfoLevel,
		log.DebugLevel,
		log.TraceLevel,
	}
}

func (hook *mongodbHook) initDatabase() {
	session, err := mgo.Dial(generateMongoConnectionURI(hook.cfg))
	hook.session = session
	if err != nil {
		panic(err)
	}

	dbName := *hook.cfg.DbName
	hook.database = session.DB(dbName)

	colName := "LogMessages"
	hook.collection = hook.database.C(colName)

	fmt.Println(dbName, colName)
}

func (hook *mongodbHook) Levels() []log.Level {
	return hook.LogLevels
}

func (hook *mongodbHook) Fire(e *log.Entry) error {
	msg := make(map[string]interface{})
	msg["TimeStamp"] = e.Time
	msg["LoggingLevel"] = e.Level
	msg["Message"] = e.Message

	err := hook.collection.Insert(msg)

	return err
}

func generateMongoConnectionURI(cfg *config.MongoDBConfig) string {
	connectionURI := "mongodb://"
	if *cfg.MongodbServerUsername != "" && *cfg.MongodbServerPassword != "" {
		connectionURI += *cfg.MongodbServerUsername + ":" + *cfg.MongodbServerPassword + "@" +
			*cfg.MongodbServerHost + ":" + *cfg.MongodbServerPort + "/" + *cfg.DbName + "?authSource=admin"
	} else {
		connectionURI += *cfg.MongodbServerHost + ":" + *cfg.MongodbServerPort
	}
	return connectionURI
}

var logger *log.Logger

func init() {
	logger = log.New()
	logger.SetOutput(ioutil.Discard)
	logger.SetFormatter(&log.TextFormatter{
		ForceColors: true,
	})
}
