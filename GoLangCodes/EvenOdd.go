package main

import "fmt"

func main() {
	fmt.Print("Enter number: ")
	var n int
	fmt.Scanln(&n)

	if n%2 == 0 {
		fmt.Println(n, "is Even")
	} else {
		fmt.Println(n, "is Odd")
	}
}