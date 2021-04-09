/*
A worker pool is a collection of goroutines working concurrently to perform a job.
In WaitGroup, we saw a collection of goroutines working concurrently but they did not have a specific job.
Once you throw channels in them, they have some job to do and becomes a worker pool.

The concept behind worker pool is maintaining a pool of worker goroutines which receives some task and returns the result.
Once they all done with their job, we collect the result. All of these goroutines use the same channel for individual purpose.
*/
package main

import (
	"fmt"
	"time"
)

// worker than make squares
func sqrWorker(tasks <-chan int, results chan<- int, instance int) { //Func sqrWorker is a worker function which takes tasks and results channel and id
	for num := range tasks {
		time.Sleep(time.Millisecond) // simulating blocking task
		fmt.Printf("[worker %v] Sending result by worker %v\n", instance, instance)
		results <- num * num
	}
}

func main() {
	fmt.Println("[main] main() started")

	tasks := make(chan int, 10)
	results := make(chan int, 10)

	// launching 3 worker goroutines
	for i := 0; i < 3; i++ {
		go sqrWorker(tasks, results, i)
	}

	// passing 5 tasks
	for i := 0; i < 5; i++ {
		tasks <- i * 2 // non-blocking as buffer capacity is 10
	}

	fmt.Println("[main] Wrote 5 tasks")

	// closing tasks
	close(tasks)

	// receving results from all workers
	for i := 0; i < 5; i++ {
		result := <-results // blocking because buffer is empty
		fmt.Println("[main] Result", i, ":", result)
	}

	fmt.Println("[main] main() stopped")
}
