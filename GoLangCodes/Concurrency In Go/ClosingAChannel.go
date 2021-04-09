/* A channel can be closed so that no more data can be sent through it.

Receiver goroutine can find out the state of the channel using val, ok := <- channel syntax
where ok is true if the channel is open or read operations can be performed and
false if the channel is closed and no more read operations can be performed.

A channel can be closed using close built-in function with syntax close(channel).

*/

package main

import (
	"fmt"
)

func greet(c chan string) {
	<-c //For John
	<-c //For Mike
}

func main() {
	fmt.Println("Main() started")

	c := make(chan string)

	go greet(c)
	c <- "John"

	close(c) //Closing of channel

	c <- "Mike"
	fmt.Println("main() stopped")
}
