package manager

import (
	"Go-Web-Dev/060_Machinery/model"
	"fmt"
	"log"

	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/backends/result"
	"github.com/RichardKnop/machinery/v1/config"
	"github.com/RichardKnop/machinery/v1/tasks"
)

// REDIS_CONFIG_URI : redis URI
const REDIS_CONFIG_URI = "redis://127.0.0.1:6379"

// TaskName : task name
const TaskName = "Introduction"

// JobManager : job manager
type JobManager struct {
	Server          *machinery.Server
	WorkerErrorChan chan error
}

// Init : initialize JobManager
func (m *JobManager) Init() {
	fmt.Println("Initialize JobManager")

	m.WorkerErrorChan = make(chan error)

	conf := &config.Config{
		Broker:        REDIS_CONFIG_URI,
		ResultBackend: REDIS_CONFIG_URI,
	}
	svr, err := machinery.NewServer(conf)
	if err != nil {
		log.Fatalln("Initialize JobManager error:", err)
	}
	m.Server = svr
	m.registerTasks()
}

func (m *JobManager) registerTasks() {
	err := m.Server.RegisterTask(TaskName, model.Hello)
	if err != nil {
		log.Fatalln("Register Tasks error:", err)
	}
}

// StartWorkerPool : generate worker pool
func (m *JobManager) StartWorkerPool(numberWorkers int) {
	worker := m.Server.NewWorker("worker", numberWorkers)
	worker.LaunchAsync(m.WorkerErrorChan)
}

// LaunchJobs : run jobs in parallel
func (m *JobManager) LaunchJobs(persons ...*model.Person) {
	fmt.Println("Launch Jobs")

	signatures := make([]*tasks.Signature, len(persons))

	for idx, p := range persons {
		signature := &tasks.Signature{
			Name: TaskName,
			Args: []tasks.Arg{
				{
					Type:  "string",
					Value: p.Name,
				},
				{
					Type:  "int",
					Value: p.Age,
				},
			},
		}
		signatures[idx] = signature
	}

	results := make([]*result.AsyncResult, len(signatures))
	for idx, task := range signatures {
		result, _ := m.Server.SendTask(task)
		results[idx] = result
	}

	for _, result := range results {
		fmt.Println(result)
	}
}
