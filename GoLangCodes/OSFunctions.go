/*The built-in os package provided by Go gives us the ability to access native operating-system features in a platform-agnostic manner.
These include creating processes, changing user permissions, file I/O and much more.
*/

/* os.PathSeparator

The os.PathSeparator constant returns the ASCII int value of the path separator.
In Unix-like systems, the path separator is / while in case of Windows, it is \
*/

package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("os.PathSeparator => %v\n", os.PathSeparator)
	fmt.Printf("os.PathSeparator => %c\n", os.PathSeparator)
}

/* os.DevNull

the os.DevNull constant returns a string specifying a null device of the operating system.
It is a mysterious virtual device that can consume data but stores nowhere, just like a black hole.
*/

package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf( "os.DevNull => %v\n", os.DevNull )
}

/* os.Args

os.Args variable returns a slice of strings that represents the arguments passed to the command
which started the program execution.

*/

package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println( "os.Args => ", os.Args )
}