package main

import (
	"fmt"
)

type Doctor struct {
	number     int
	actorName  string
	companions []string //slices of string
}

func main() {
	aDoctor := Doctor{
		number:    3,
		actorName: "ABC",
		companions: []string{
			"def",
			"ghi",
			"jkl",
		},
	}
	//fmt.Println(aDoctor)
	fmt.Println(aDoctor.actorName) //importing a data element from struct
}
