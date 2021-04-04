package main

import (
	"fmt"
)

func main() {
	a := [3]int{1, 2, 3}
	b := &a[0]
	c := &a[1]
	fmt.Printf("%v %p %p\n", a, b, c)
	/* Beyond this no arithmetic operation can be done like C and C++
	So do them in run time import package "unsafe" if you wanna use the operations.
	*/
}
