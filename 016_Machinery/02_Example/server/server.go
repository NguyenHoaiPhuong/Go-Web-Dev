package server

import (
	"log"

	"github.com/NguyenHoaiPhuong/Go-Web-Dev/016_Machinery/02_Example/task"
	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
)

var server *machinery.Server
var errorsChan chan<- error

// GetServer : return singleton server
func GetServer() *machinery.Server {
	if server == nil {
		var cnf = &config.Config{
			// Broker:       RABBITMQ_CONF_URI,
			Broker:          "redis://localhost:6380",
			DefaultQueue:    "machinery_tasks",
			ResultBackend:   "'mongodb://localhost:27017",
			ResultsExpireIn: 3600000,
		}
		var err error
		server, err = machinery.NewServer(cnf)
		if err != nil {
			log.Fatalln("Create new server error:", err)
		}
	}
	return server
}

// InitServer : init machinery server
func InitServer(configPath string) {
	cnf, err := config.NewFromYaml(configPath, true)
	server, err = machinery.NewServer(cnf)
	if err != nil {
		log.Fatalln("Create new server error:", err)
	}
}

// LaunchWorkerPool : launch workers
func LaunchWorkerPool(workerName string, numWorkers int) {
	worker := server.NewWorker(workerName, numWorkers)
	errorsChan = make(chan error)
	worker.LaunchAsync(errorsChan)
	// err := worker.Launch()
	// if err != nil {
	// 	log.Fatalln("Launch Worker Pool error:", err)
	// }
}

// RegisterTasks : register ADD and MULTIPLY tasks
func RegisterTasks() {
	// We can register task one by one
	// server.RegisterTask(task.ADD, task.Add)
	// server.RegisterTask(task.MULTIPLY, task.Multiply)

	// Or we can register multiple tasks once
	server.RegisterTasks(map[string]interface{}{
		task.ADD:      task.Add,
		task.MULTIPLY: task.Multiply,
	})
}
