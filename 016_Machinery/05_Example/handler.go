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
	for idx, param := range taskParams {
		signature := createTaskSignature(param)
		signatures[idx] = signature
	}

	chain, err := tasks.NewChain(signatures...)
	checkError(err)

	chainAsyncResult, err := srv.SendChain(chain)
	checkError(err)

	log.Println("chordAsyncResult:", chainAsyncResult)
	c.JSON(http.StatusOK, gin.H{"Status": "In progress", "Job": chainAsyncResult})
}

func createTaskSignature(param *TaskParam) *tasks.Signature {
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

	return signature
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
