/*
A map maps keys to values.

The zero value of a map is nil. A nil map has no keys, nor can keys be added.

The make function returns a map of the given type, initialized and ready for use.
*/

package main

import (
	"fmt"
)

func main() {
	statePopulations := map[string]int{
		"ABC": 23232,
		"DEF": 23215,
		"GHI": 96922,
		"JKL": 23421,
		"MNO": 86963,
	}
	fmt.Println(statePopulations)
}
