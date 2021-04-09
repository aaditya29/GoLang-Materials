/*channels are first-class values and can be used anywhere like other values:
as struct elements, function arguments, returning values and even like a type for another channel.
Here, we are interested in using a channel as the data type of another channel.
*/

package main

import "fmt"

// gets a channel and prints the greeting by reading from channel
func greet(c chan string) {
	fmt.Println("Hello " + <-c + "!")
}

// gets a channels and writes a channel to it
func greeter(cc chan chan string) {
	c := make(chan string)
	cc <- c
}

func main() {
	fmt.Println("main() started")

	// make a channel `cc` of data type channel of string data type
	cc := make(chan chan string)

	go greeter(cc) // start `greeter` goroutine using `cc` channel

	// receive a channel `c` from `greeter` goroutine
	c := <-cc

	go greet(c) // start `greet` goroutine using `c` channel

	// send data to `c` channel
	c <- "John"

	fmt.Println("main() stopped")
}
