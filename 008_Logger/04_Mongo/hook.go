package main

import (
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

	dbName := *hook.cfg.ConfigDbName
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
			*cfg.MongodbServerHost + ":" + *cfg.MongodbServerPort + "/" + *cfg.ConfigDbName + "?authSource=admin"
	} else {
		connectionURI += *cfg.MongodbServerHost + ":" + *cfg.MongodbServerPort
	}
	return connectionURI
}
