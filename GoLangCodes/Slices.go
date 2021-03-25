package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var s []int = make([]int, 3)
	var in string
	fmt.Println("Please enter an integer or E to exit:\n")
	for true {
		fmt.Scanln(&in)
		if in == "E" {
			break
		}
		ap, err := strconv.Atoi(in)
		if err != nil {
			fmt.Println("Provided input is not correct")
			continue
		}
		s = append(s, ap)
		sort.Ints(s[:])
		fmt.Println(s)
		fmt.Println("Enter an integer again or E to exit: \n")
	}
}
