package main

import (
	"fmt"
)

func main() {
	var a int = 42
	var b *int = &a //b is pointer to an integer and wanna point to a
	fmt.Println(a, b)
}
