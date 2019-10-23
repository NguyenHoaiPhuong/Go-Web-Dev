package config

import (
	"fmt"

	"github.com/paked/configure"
)

// MongoDBConfig includes all configs related to mongodb settings
type MongoDBConfig struct {
	MongodbServerHost     *string
	MongodbServerPort     *string
	MongodbServerUsername *string
	MongodbServerPassword *string
	DbName                *string
}

// Copy : copy values of all fields in the existing config and return a new config
func (cfg *MongoDBConfig) Copy() *MongoDBConfig {
	newCfg := &MongoDBConfig{
		MongodbServerHost:     new(string),
		MongodbServerPort:     new(string),
		MongodbServerUsername: new(string),
		MongodbServerPassword: new(string),
		DbName:                new(string),
	}
	*newCfg.MongodbServerHost = *cfg.MongodbServerHost
	*newCfg.MongodbServerPort = *cfg.MongodbServerPort
	*newCfg.MongodbServerUsername = *cfg.MongodbServerUsername
	*newCfg.MongodbServerPassword = *cfg.MongodbServerPassword
	*newCfg.DbName = *cfg.DbName

	return newCfg
}

// Config includes all configs for running the app
type Config struct {
	LoggingLevel *int
	*MongoDBConfig
}

// newConfig return default configs
func newConfig() (*configure.Configure, *Config) {
	var conf = configure.New()
	var cfg = Config{
		MongoDBConfig: &MongoDBConfig{
			MongodbServerHost:     conf.String("mongodbServerHost", "localhost", "The host address with which the mongodb instance containing the network setting database will be reachable"),
			MongodbServerPort:     conf.String("mongodbServerPort", "27017", "The host port with which the mongodb instance containing the network setting database will be reachable"),
			MongodbServerUsername: conf.String("mongodbServerUsername", "", "mongodb username"),
			MongodbServerPassword: conf.String("mongodbServerPassword", "", "mongodb password"),
			DbName:                conf.String("dbName", "test", "The name of the database that contains the description of the network to simulate, default to 'network'"),
		},
		LoggingLevel: conf.Int("loggingLevel", 4, "Integer starts from 0 to 6 represented for each logging level: 0 = PanicLevel,	1 = FatalLevel, 2 = ErrorLevel,	3 = WarnLevel, 4 = InfoLevel, 5 = DebugLevel, 6 = TraceLevel. Default is Info level."),
	}
	return conf, &cfg
}

// SetupAndGetConfig parses configs set in global env, json file and command line
func SetupAndGetConfig(jsonPath string) *Config {
	var conf, cfg = newConfig()

	// Reverse order, Flag is the most important here
	conf.Use(configure.NewFlag())
	conf.Use(configure.NewJSONFromFile(jsonPath))
	conf.Use(configure.NewEnvironment())

	conf.Parse()

	cfg.print()

	return cfg
}

func (cfg *Config) print() {
	fmt.Println("MongodbServerHost :", *cfg.MongodbServerHost)
	fmt.Println("MongodbServerPort :", *cfg.MongodbServerPort)
	fmt.Println("MongodbServerUsername :", *cfg.MongodbServerUsername)
	fmt.Println("MongodbServerPassword :", *cfg.MongodbServerPassword)
	fmt.Println("DbName :", *cfg.DbName)
	fmt.Println("LoggingLevel :", *cfg.LoggingLevel)
}
