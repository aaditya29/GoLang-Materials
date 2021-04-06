/*
Variadic functions can be called with any number of trailing arguments.
For example, fmt.Println is a common variadic function.
*/

package main

import (
	"fmt"
)

func main() {
	sum(1, 2, 3, 4, 5)
}

func sum(values ...int) /*making variadic function followed by its type*/ {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Println("The sum is ", result)
}
