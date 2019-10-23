package main

import (
	"os"
	"time"

	"github.com/NguyenHoaiPhuong/Go-Web-Dev/008_Logger/04_Mongo/config"

	"github.com/globalsign/mgo"
	log "github.com/sirupsen/logrus"
)

type mongodbHook struct {
	cfg        *config.MongoDBConfig
	session    *mgo.Session
	database   *mgo.Database
	collection *mgo.Collection
}

func (hook *mongodbHook) Init(cfg *config.MongoDBConfig) {
	hook.cfg = cfg.Copy()
	hook.initDatabase()
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
}

func (mongodbHook) Levels() []log.Level {
	return []log.Level{
		log.PanicLevel,
		log.FatalLevel,
		log.ErrorLevel,
		log.WarnLevel,
		log.InfoLevel,
		log.DebugLevel,
		log.TraceLevel,
	}
}

func (hook mongodbHook) Fire(e *log.Entry) error {
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
	logger.SetOutput(os.Stdout)
	logger.SetLevel(log.ErrorLevel)
}

func main() {
	cfg := config.SetupAndGetConfig("./resources/conf.json")

	entry := logger.WithField("function", "main")
	entry.Info("Start ...")

	hook := new(mongodbHook)
	hook.Init(cfg.MongoDBConfig)
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
