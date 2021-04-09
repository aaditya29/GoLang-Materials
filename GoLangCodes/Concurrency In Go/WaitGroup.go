/*
WaitGroup is a struct with a counter value which tracks
how many goroutines were spawned and how many have completed their job.
This counter when reaches zero, means all goroutines have done their job.
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

func service(wg *sync.WaitGroup, instance int) {
	time.Sleep(2 * time.Second)
	fmt.Println("Service called on instance", instance)
	wg.Done() //Decrement counter
}

func main() {
	fmt.Println("main() started")
	var wg sync.WaitGroup //Creating empty struct(with zero-value filerds) wg of type sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1) //Incremenet counter
		go service(&wg, i)
	}

	wg.Wait() //Blocking here
	fmt.Println("main() stopped.")

	/*
		after creating wg, we ran for loop for 3 times.
		In each turn, we launched 1 goroutine and incremented the counter by 1.
		That means, now we have 3 goroutines waiting to be executed and WaitGroup counter is 3.

		we passed a pointer to wg in goroutine.
		This is because in goroutine, once we are done with whatever the goroutine was supposed to do,
		we need to call Done method to decrement the counter.
		If wg was passed as a value, wg in main would not get decremented.

		After for loop has done executing, we still did not pass control to other goroutines.
		This is done by calling Wait method on wg like wg.Wait().
		This will block the main goroutine until the counter reaches 0.
		Once the counter reaches 0 because from 3 goroutines, we called Done method on wg 3 times, main goroutine will unblock and starts executing further code.
	*/
}
