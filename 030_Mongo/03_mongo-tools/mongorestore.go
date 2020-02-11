package main

// mongorestore $MONGO_BACKUP_DATA_ROOT_FOLDER_PATH/$RESTORED_DB_NAME --host ${DB_HOST} --port ${DB_PORT} --drop --db ${RESTORED_DB_NAME} ---username ${DB_USER} --password ${DB_PASS} --authenticationDatabase admin

import (
	// "github.com/mongodb/mongo-tools-common/common/db"

	"log"

	"github.com/mongodb/mongo-tools-common/db"
	commonOpts "github.com/mongodb/mongo-tools-common/options"
	"github.com/mongodb/mongo-tools/mongorestore"
)

const (
	MONGO_BACKUP_DATA_ROOT_FOLDER_PATH = "./db"
	DB_HOST                            = "localhost"
	DB_PORT                            = "27018"
	DB_USER                            = "simcel"
	DB_PASS                            = "simcel"
	RESTORED_DB_NAME                   = "ITV"
)

// InitMongorestore : initialize mongorestore with given options
func InitMongorestore(serverHost, serverPort, authUser, authPass, restoredDBame, localPath string) *mongorestore.MongoRestore {
	var toolOptions *commonOpts.ToolOptions
	auth := &commonOpts.Auth{
		Username: authUser,
		Password: authPass,
		Source:   "admin",
	}
	connection := &commonOpts.Connection{
		Host: serverHost,
		Port: serverPort,
	}
	nameSpace := &commonOpts.Namespace{DB: restoredDBame}
	toolOptions = &commonOpts.ToolOptions{
		Connection: connection,
		Auth:       auth,
		Namespace:  nameSpace,
	}

	inputOpts := &mongorestore.InputOptions{}
	outputOpts := &mongorestore.OutputOptions{
		NumParallelCollections: 1,
		NumInsertionWorkers:    1,
		WriteConcern:           "majority",
	}

	nsOpts := &mongorestore.NSOptions{
		DB: restoredDBame,
	}

	dbPath := localPath + "/" + restoredDBame

	provider, err := db.NewSessionProvider(*toolOptions)
	checkError(err)

	return &mongorestore.MongoRestore{
		ToolOptions:     toolOptions,
		OutputOptions:   outputOpts,
		InputOptions:    inputOpts,
		NSOptions:       nsOpts,
		SessionProvider: provider,
		TargetDirectory: dbPath,
	}
}

func main() {
	mr := InitMongorestore(DB_HOST, DB_PORT, DB_USER, DB_PASS, RESTORED_DB_NAME, MONGO_BACKUP_DATA_ROOT_FOLDER_PATH)

	// log.Println("Auth:", mr.ToolOptions.Auth)
	// log.Println("Connection:", mr.ToolOptions.Connection)
	// log.Println("DBPath:", mr.ToolOptions.DBPath)
	// log.Println("DB:", mr.ToolOptions.DB)
	// log.Println("Username:", mr.ToolOptions.Username)
	// log.Println("Password:", mr.ToolOptions.Password)

	// err := mr.Restore()
	// checkError(err)
	result := mr.Restore()
	log.Println(result)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
