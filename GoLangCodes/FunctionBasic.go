package main

import (
	"fmt"
)

func main() {
	sayMessage("Hello Go!") //Function calling
}

func sayMessage(msg string) { //function declaration
	fmt.Println(msg)
}
