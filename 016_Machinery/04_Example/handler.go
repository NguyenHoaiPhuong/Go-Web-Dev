package main

import (
	"log"
	"net/http"

	"github.com/NguyenHoaiPhuong/Go-Web-Dev/016_Machinery/02_Example/task"
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

	addSignatures := make([]*tasks.Signature, 0)
	mulSignatures := make([]*tasks.Signature, 0)
	for _, param := range taskParams {
		switch param.TaskName {
		case task.ADD:
			signature := createTaskSignature(param)
			addSignatures = append(addSignatures, signature)
		case task.MULTIPLY:
			signature := createTaskSignature(param)
			mulSignatures = append(mulSignatures, signature)
		default:
			log.Fatalln("Unknown task name", param.TaskName)
		}

	}

	// Group of addition tasks
	group, err := tasks.NewGroup(addSignatures...)
	checkError(err)

	// Chord
	if len(mulSignatures) > 1 {
		log.Fatalln("Please input only 1 multiplying task")
	}
	chord, err := tasks.NewChord(group, mulSignatures[0])
	checkError(err)

	//The second parameter of SendChord specifies the number of concurrent sending tasks. 0 means unlimited.
	chordAsyncResult, err := srv.SendChord(chord, 0)
	checkError(err)

	log.Println("chordAsyncResult:", chordAsyncResult)
	c.JSON(http.StatusOK, gin.H{"Status": "In progress", "Job": chordAsyncResult})
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
