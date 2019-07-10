package manager

import (
	"Go-Web-Dev/060_Machinery/model"
	"fmt"
	"log"

	"github.com/RichardKnop/machinery/v1"
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
func (m *JobManager) LaunchJobs(persons ...*model.Person) []string {
	fmt.Println("Launch Jobs")

	strResults := make([]string, 0)

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

	// group, err := tasks.NewGroup(signatures...)
	// if err != nil {
	// 	log.Fatalln("Creating new group error:", err)
	// }
	// results, err := m.Server.SendGroup(group, 8)
	// if err != nil {
	// 	log.Fatalln("Send group of tasks error:", err)
	// }

	// for _, result := range results {
	// 	state := result.GetState()
	// 	strRes := state.Results[0].Value.(string)
	// 	strResults = append(strResults, strRes)
	// 	fmt.Println(state)
	// }

	for _, task := range signatures {
		result, err := m.Server.SendTask(task)
		if err != nil {
			log.Fatalln("Send task error:", err)
		}
		state := result.GetState()
		if len(state.Results) > 0 {
			strRes := state.Results[0].Value.(string)
			strResults = append(strResults, strRes)
		}
	}

	return strResults
}
