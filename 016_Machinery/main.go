package main

import (
	"Go-Web-Dev/060_Machinery/manager"
	"Go-Web-Dev/060_Machinery/model"
	"fmt"
)

func main() {
	akagi := &model.Person{
		Name: "Nguyen Hoai Phuong",
		Age:  33,
	}
	yushin := &model.Person{
		Name: "Binh Vo",
		Age:  27,
	}
	mogami := &model.Person{
		Name: "Tien Nguyen",
		Age:  26,
	}

	jobManager := new(manager.JobManager)
	jobManager.Init()
	jobManager.StartWorkerPool(8)
	results := jobManager.LaunchJobs(akagi, yushin, mogami)
	for _, result := range results {
		fmt.Println(result)
	}
	fmt.Println("Finish")
}
