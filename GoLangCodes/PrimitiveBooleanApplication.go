package main

import (
	"fmt"
)

func main() {
	n := 1 == 1
	m := 1 == 2
	fmt.Printf("%v and data type is %T\n", n, n)
	fmt.Printf("%v and data type is %T", m, m)
}
