package main

import "fmt"

func main() {
	a, b, c := 10, 20, 30

	fmt.Println(
		"(a + b = c) :", a, "+", b, "=", c,
	)

	fmt.Printf(
		"(a + b = c) : %d + %d = %d\n", a, b, c,
	)
}
