/* By default, a channel buffer size is 0 also called as unbuffered channel.
Whatever written to the channel is immediately available to read.

When the buffer size is non-zero n, goroutine is not blocked until after buffer is full.
When the buffer is full, any value sent to the channel is added to the buffer by throwing out last value
in the buffer which is available to read (where the goroutine will be blocked).
But there is a catch, read operation on buffered is thirsty. That means, once read operation starts,
it will continue until the buffer is empty.
Technically, that means goroutine that reads buffer channel will not block until the buffer is empty.

DEFINING A BUFFERED CHANNEL:
c := make(chan Type, n)
This will create a channel of a data type Type with buffer size n. Until the channel receives n+1 send operations, it won’t block the current goroutine.
*/

package main

import (
	"fmt"
)

func squares(c chan int) {
	for i := 0; i <= 3; i++ {
		num := <-c
		fmt.Println(num + num)
	}
}

func main() {
	fmt.Println("main() started ")
	c := make(chan int, 3) //channel c having buffer capacity of 3

	go squares(c)

	c <- 1
	c <- 2
	c <- 3

	fmt.Println("main() stopped")
}
