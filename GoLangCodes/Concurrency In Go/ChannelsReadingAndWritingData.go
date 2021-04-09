/* Go provide very easy to remember left arrow syntax <- to read and write data from a channel:
c <- data

Above syntax means, we want to push or write data to the channel c.
*/

package main

import (
	"fmt"
)

func greet(c chan string) { //declaring greet function which accepts a channel c of data type string
	fmt.Println("Hello " + <-c + "!")
}

func main() {
	fmt.Println("main() started")
	c := make(chan string) //Creating channel c of type string using make function

	go greet(c) //Passing channel c to the greet function but executing as goroutine

	c <- "John"
	fmt.Println("main() stopped")
}
