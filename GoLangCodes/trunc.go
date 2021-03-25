//WAP which prompts the user to enter a floating point number and prints the integer which is a truncated version of the floating point number that was entered.
package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Print("Please enter a floating point number=> \n")
	fmt.Scanln(&number)
	fmt.Print("\nConverted Integer is=> \n", number)
}
