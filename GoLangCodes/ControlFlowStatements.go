//Simple If Statement
package main

import (
	"fmt"
	"math"
)

/*
func main() {
	if true {
		fmt.Println("The test is true")
	}
}
*/
/*
func main() {
	number := 50
	guess := 30
	if guess < number {
		fmt.Println("Too low")
	}
	if guess > number {
		fmt.Println("Too high")
	}
	if guess == number {
		fmt.Println("You got it!")
	}
	fmt.Println(number <= guess, number >= guess, number != guess)
}
*/

//If And Else Statements:

func main() {
	myNum := 0.1
	if myNum == math.Pow(math.Sqrt(myNum), 2) {
		fmt.Println("These are the same")
	} else {
		fmt.Println("These are different.\n")
	}

}
