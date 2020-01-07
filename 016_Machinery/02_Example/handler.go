package main

import (
	"log"
	"net/http"

	"github.com/NguyenHoaiPhuong/Go-Web-Dev/016_Machinery/02_Example/server"
	"github.com/RichardKnop/machinery/v1/tasks"
	"github.com/gin-gonic/gin"
)

// TaskParams struct
type TaskParams struct {
	TaskName string
	Args     []int64
}

// CreateTask : create a single task for master
func CreateTask(c *gin.Context) {
	srv := server.GetServer()

	var param TaskParams
	c.BindJSON(&param)

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

	asyncResult, err := srv.SendTask(signature)
	if err != nil {
		panic(err)
	}
	log.Println("asyncResult:", asyncResult)

	c.JSON(http.StatusOK, gin.H{"Status": "In progress", "Job": asyncResult})
}
