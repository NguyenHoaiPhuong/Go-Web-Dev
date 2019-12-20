package main

import (
	"fmt"

	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
)

func main() {
	cnf, err := config.NewFromYaml("config.yml", true)
	server, err := machinery.NewServer(cnf)
	if err != nil {
		fmt.Println("Create new server error:", err)
		panic(err)
	}

	worker := server.NewWorker("worker_name", 10)
	err = worker.Launch()
	if err != nil {
		fmt.Println("Launch worker error:", err)
		panic(err)
	}

	// We can register task one by one
	server.RegisterTask("add", Add)
	server.RegisterTask("multiply", Multiply)

	// Or we can register multiple tasks once
	server.RegisterTasks(map[string]interface{}{
		"add":      Add,
		"multiply": Multiply,
	})

}
