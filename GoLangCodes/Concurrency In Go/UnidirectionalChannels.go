/* Creating channels which are unidirectional i.e.
 receive-only channels which allow only read operation on them and send-only
 channels which allow only to write operation on them.

The unidirectional channel is also created using make function but with additional arrow syntax.

roc := make(<-chan int)
soc := make(chan<- int)

Above roc is receive-only channel as arrow direction in make function points away from chan keyword.
While soc is send-only channel where arrow direction in make function points towards chan keyword.
*/

package main

import (
	"fmt"
)

func greet(roc <-chan string) {
	fmt.Println("Hello " + <-roc + "!")
}

func main() {
	fmt.Println("main() started")
	c := make(chan string)

	go greet(c)

	c <- "John"
	fmt.Println("main() stopped")
}
