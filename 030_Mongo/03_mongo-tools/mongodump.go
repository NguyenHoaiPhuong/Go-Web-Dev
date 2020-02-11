package main

import (
	"github.com/mongodb/mongo-tools-common/log"
	commonOpts "github.com/mongodb/mongo-tools-common/options"
	"github.com/mongodb/mongo-tools/mongodump"
)

// mongodump --out $MONGO_BACKUP_DATA_ROOT_FOLDER_PATH/ --host ${DB_HOST} --port ${DB_PORT} --db ${SAVED_DB_NAME} --username ${DB_USER} --password ${DB_PASS} --authenticationDatabase admin

const (
	MONGO_BACKUP_DATA_ROOT_FOLDER_PATH = "./db"
	DB_HOST                            = "localhost"
	DB_PORT                            = "27018"
	DB_USER                            = "simcel"
	DB_PASS                            = "simcel"
	SAVED_DB_NAME                      = "ITV"
)

// InitMongodump : initialize mongodump with given options
func InitMongodump(serverHost, serverPort, authUser, authPass, savedDBName, outPath string) *mongodump.MongoDump {
	var toolOptions *commonOpts.ToolOptions
	ssl := &commonOpts.SSL{
		UseSSL: false,
	}
	auth := &commonOpts.Auth{
		Username: authUser,
		Password: authPass,
		Source:   "admin",
	}
	namespace := &commonOpts.Namespace{
		DB: savedDBName,
	}
	connection := &commonOpts.Connection{
		Host: serverHost,
		Port: serverPort,
	}
	toolOptions = &commonOpts.ToolOptions{
		SSL:        ssl,
		Namespace:  namespace,
		Connection: connection,
		Auth:       auth,
		Verbosity:  &commonOpts.Verbosity{},
		URI: &commonOpts.URI{
			ConnectionString: "",
		},
	}

	outputOptions := &mongodump.OutputOptions{
		NumParallelCollections: 1,
		Out:                    outPath,
	}
	inputOptions := &mongodump.InputOptions{
		Query: "",
	}

	log.SetVerbosity(toolOptions.Verbosity)

	return &mongodump.MongoDump{
		ToolOptions:   toolOptions,
		InputOptions:  inputOptions,
		OutputOptions: outputOptions,
	}
}

func main() {
	md := InitMongodump(DB_HOST, DB_PORT, DB_USER, DB_PASS, SAVED_DB_NAME, MONGO_BACKUP_DATA_ROOT_FOLDER_PATH)

	err := md.Init()
	checkError(err)
	err = md.Dump()
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
