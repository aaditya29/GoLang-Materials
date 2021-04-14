//Monitoring multiple goroutine servers
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func main() {
	//Creating a WaitGroup which montors goroutines
	wg := new(sync.WaitGroup)

	//Adding two goroutines to 'wg' WaitGroup
	wg.Add(2)

	//Creating a default route handler
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Hello: "+req.Host)
	})

	//goroutine to launch a server on port 9000
	go func() {
		log.Fatal(http.ListenAndServe(":9000", nil))
		wg.Done() //First go routine finished
	}()

	//goroutine to launch a server on port 9001
	go func() {
		log.Fatal(http.ListenAndServe(":9001", nil))
		wg.Done() //Second goroutine finished
	}()

	wg.Wait() //Wait until WaitGroup is done ie.e main method would not exit until the counter in the WaitGroup reaches zero
}
