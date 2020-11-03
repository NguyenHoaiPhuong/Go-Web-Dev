package main

import (
	"fmt"
	"time"
)

const (
	numJobs    = 5
	numWorkers = 3
)

func worker(id int, jobs <-chan int, results chan int) {
	for job := range jobs {
		fmt.Println("Worker", id, "started job", job)
		time.Sleep(time.Second)
		fmt.Println("Worker", id, "finished job", job)
		results <- job
	}
}

func createJobs(jobs chan int) {
	for i := 1; i <= numJobs; i++ {
		jobs <- i
	}
}

func main() {
	jobs := make(chan int)
	results := make(chan int, numJobs)

	for i := 1; i <= numWorkers; i++ {
		go worker(i, jobs, results)
	}

	// create jobs
	createJobs(jobs)
	// close(jobs)

	// results
	for i := 0; i < numJobs; i++ {
		fmt.Println(<-results)
	}
}
