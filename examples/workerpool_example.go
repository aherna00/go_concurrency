package examples

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job)
		time.Sleep(time.Second) // Simulate work
		fmt.Printf("Worker %d finished job %d\n", id, job)
		results <- job * 2 // Send result back
	}
}

func startWork() {
	const numJobs = 500000 //large amount of data to be processed
	const numWorkers = 10  //process data with 10 go routines

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= numWorkers; w++ { //spawn 10 go routines
		go worker(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ { //start handing out jobs to go routines, until all jobs finished
		jobs <- j
	}

	close(jobs)

	for a := 1; a <= numJobs; a++ { //aggregate results
		result := <-results
		fmt.Printf("Result: %d\n", result)
	}
}
