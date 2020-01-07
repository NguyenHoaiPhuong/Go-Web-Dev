package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/NguyenHoaiPhuong/Go-Web-Dev/016_Machinery/02_Example/server"
	"github.com/gin-gonic/gin"

	"github.com/RichardKnop/machinery/v1/tasks"
)

var mode string

func main() {
	flag.StringVar(&mode, "mode", "master", "Default mode is master. Another available mode is worker")
	flag.Parse()
	log.Println("Running Mode:", mode)

	server.InitServer("./config/config.yml")

	switch mode {
	case "master":
		router := gin.Default()
		router.RedirectTrailingSlash = true
		router.RedirectFixedPath = true
		router.POST("/tasks", CreateTask)
		srv := &http.Server{
			Addr:    ":9000",
			Handler: router,
		}
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	case "worker":
		server.LaunchWorkerPool("yushin", 10)
		server.RegisterTasks()
		// Nothing in the startwroker pool is blocking. Need to wait to avoid ending the process.
		for {
			time.Sleep(1000 * time.Millisecond)
		}
	default:
		log.Fatalln("Unknown running mode")
	}

}

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
