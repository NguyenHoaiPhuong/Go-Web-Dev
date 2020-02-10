package main

import (
	commonOpts "github.com/mongodb/mongo-tools/common/options"
	"github.com/mongodb/mongo-tools/mongodump"
	"github.com/mongodb/mongo-tools/mongodump/options"
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
func InitMongodump(serverHost, serverPort, authUser, authPass, savedDBame, outPath string) *mongodump.MongoDump {
	var toolOptions *commonOpts.ToolOptions
	ssl := &commonOpts.SSL{
		UseSSL: false,
	}
	auth := &commonOpts.Auth{
		Username: authUser,
		Password: authPass,
		Source:   "admin",
	}
	connection := &commonOpts.Connection{
		Host: serverHost,
		Port: serverPort,
	}
	toolOptions = &commonOpts.ToolOptions{
		SSL:        ssl,
		Connection: connection,
		Auth:       auth,
	}
	toolOptions.Namespace = &commonOpts.Namespace{DB: savedDBame}

	outputOptions := &options.OutputOptions{
		Out: outPath,
	}

	inputOptions := &options.InputOptions{}

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
