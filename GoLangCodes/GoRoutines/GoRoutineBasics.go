/* GO ROUTINE

goroutine is simply a function or method that is running in background concurrently with other goroutines.
It’s not a function or method definition that determines if it is a goroutine, it is determined by how we call it.

Go provides a special keyword go to create a goroutine. When we call a function or a method with go prefix,
that function or method executes in a goroutine.
*/

package main

import (
	"fmt"
	"time"
)

//Without go routine
func printHello() {
	fmt.Println("Hello World")
}

/*
func main() {
	fmt.Println("Main execution started")

	printHello() //calling function

	fmt.Println("Main execution stopped")
}
*/

//Now creating GoRoutine

/*
	func main() {
	fmt.Println("Main execution started")

	go printHello() //creating goroutine
	fmt.Println("Main execution stopped")
*/

/* OUTPUT:
main execution started
main execution stopped

Here no Hello World is printed. Reason?

goroutines always run in the background. When a goroutine is executed, here, Go does not block the program execution,
unlike normal function call as we have seen in the previous example.
Instead, the control is returned immediately to the next line of code and any returned value from goroutine is ignored.
But even then, why we can’t see the function output?

By default, every Go standalone application creates one goroutine.
This is known as the main goroutine that the main function operates on.
In the above case, the main goroutine spawns another goroutine of printHello function, let’s call it printHello goroutine.
Hence when we execute the above program, there are two goroutines running concurrently.
As we saw in the earlier program, goroutines are scheduled cooperatively.

Hence when the main goroutine starts executing, go scheduler dot not pass control to the printHello goroutine until
the main goroutine does not execute completely. Unfortunately,
when the main goroutine is done with execution, the program terminates immediately and scheduler did not get time to schedule printHello goroutine.

But using blocking conditions i.e. time.Sleep() we can pass control to other goroutines.
*/

func main() {
	fmt.Println("Main execution started")

	go printHello() //Creating goroutine

	time.Sleep(10 * time.Millisecond) //Scheduling another goroutine
	fmt.Println("Main exeuction stopped")
}
