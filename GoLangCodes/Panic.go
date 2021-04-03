//Raises exception
package main

import "fmt"

func main() {
	fmt.Println("start")
	panic("Something Bad Happened")
	fmt.Println("End")
}
