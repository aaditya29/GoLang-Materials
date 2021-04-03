/*A defer statement pushes a function call onto a list.
The list of saved calls is executed after the surrounding function returns.
Defer is commonly used to simplify functions that perform various clean-up actions.
*/

package main

import (
	"fmt"
)

//Before defer
/*
func main() {
	fmt.Println("Start")
	fmt.Println("Middle")
	fmt.Println("End")
}
*/

//Using defer
func main() {
	fmt.Println("Start")
	defer fmt.Println("Middle")
	fmt.Println("End")
}
