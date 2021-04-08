//Creating two goroutines from two function calls.
//One function prints char of the string and another prints digit of integer
package main

import (
	"fmt"
	"time"
)

func getChars(s string) {
	for _, c := range s {
		fmt.Printf("%c", c)
	}
}

func getDigits(s []int) {
	for _, d := range s {
		fmt.Printf("%d", d)
	}
}

func main() {
	fmt.Println("Main execution started.")

	go getChars("Hello") //Creating goroutine of getChars

	go getDigits([]int{1, 2, 3, 4, 5}) //Creating goroutine of getDigits

	//Scheduling another goroutine
	time.Sleep(time.Millisecond)

	fmt.Println("\nMain exeuction stopped.")
}
