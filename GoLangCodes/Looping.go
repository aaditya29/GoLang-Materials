package main

//Simple loop
/*
func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}
*/
import (
	"fmt"
)

func main() {
	s := "hello world!"
	for k, v := range s {
		fmt.Println(k, string(v))
	}
}
