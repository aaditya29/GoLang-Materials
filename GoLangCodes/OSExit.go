//Demonstrating working of os.Exit function

/*The os.Exit(code int) function kills the current process and exits the program with a code integer status code.

The status code0 indicates that the process was exited successfully while
the status code 1 indicates the process was exited with a general error.

*/

package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("main() started!")

	// defer function execution
	defer func() {
		fmt.Println("main completed!")
	}()

	// run function after 2 seconds
	time.AfterFunc(2*time.Second, func() {
		os.Exit(1) // exit with status code : 1
	})

	// sleep goroutine for 3 seconds
	time.Sleep(3 * time.Second)

}
