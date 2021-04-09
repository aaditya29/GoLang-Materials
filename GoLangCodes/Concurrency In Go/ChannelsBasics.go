/*
A channel is a communication object using which goroutines can communicate with each other.
Technically, a channel is a data transfer pipe where data can be passed into or read from.
Hence one goroutine can send data into a channel, while other goroutines can read that data from the same channel.
*/

/*
Go provides chan keyword to create a channel.
A channel can transport data of only one data type. No other data types are allowed to be transported from that channel.
*/

package main

import (
	"fmt"
)

/*
func main() {

	var c chan int //Declaring a channel c which can transport data type of int
	fmt.Println(c)
	//OUTPUT: <nil>
	//You can not pass data to or read data from a channel which is nil. Hence
	//we have to use make function to create a ready-to-use channel.


}
*/

//Creating ready-to-use channel with help of make function

func main() {
	c := make(chan int) //

	fmt.Printf("Type of 'c' is %T\n", c)
	fmt.Printf("Value of 'c' is %v\n", c)
	/* OUTPUT:
	type of `c` is chan int
	value of `c` is 0xc00009e000

	Value of channel 'c' is a memory address which shows Channels by default are pointers.
	*/
}
