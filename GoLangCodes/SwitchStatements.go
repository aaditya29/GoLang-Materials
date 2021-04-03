package main

import (
	"fmt"
)

/*
func main() {
	switch 2 {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("not one or two.")
	}
}
*/

//Alternate way to write
func main() {
	i := 10
	switch {
	case i <= 10:
		fmt.Println("Less that or equal to ten.")
	case i <= 20:
		fmt.Println("Less than or equal to twenty.")
	default:
		fmt.Println("Greater than twenty.")
	}
}
