package main

import (
	"log"
	"net/http"

	"github.com/NguyenHoaiPhuong/Go-Web-Dev/016_Machinery/03_Example/server"
	"github.com/RichardKnop/machinery/v1/tasks"
	"github.com/gin-gonic/gin"
)

// TaskParam struct
type TaskParam struct {
	TaskName string
	Args     []int64
}

// TaskParams : slice of TaskParam
type TaskParams []*TaskParam

// CreateTask : create a single task for master
func CreateTask(c *gin.Context) {
	srv := server.GetServer()

	taskParams := make(TaskParams, 0)
	c.BindJSON(&taskParams)

	signatures := make([]*tasks.Signature, len(taskParams))
	for i, param := range taskParams {
		args := make([]tasks.Arg, len(param.Args))
		for idx, arg := range param.Args {
			args[idx] = tasks.Arg{
				Type:  "int64",
				Value: arg,
			}
		}
		signature := &tasks.Signature{
			Name:       param.TaskName,
			Args:       args,
			RetryCount: 3,
		}
		signatures[i] = signature
	}

	group, err := tasks.NewGroup(signatures...)
	checkError(err)
	asyncResults, err := srv.SendGroup(group, 0) //The second parameter specifies the number of concurrent sending tasks. 0 means unlimited.
	checkError(err)

	for _, asyncResult := range asyncResults {
		log.Println("asyncResult:", asyncResult)
		c.JSON(http.StatusOK, gin.H{"Status": "In progress", "Job": asyncResult})
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
