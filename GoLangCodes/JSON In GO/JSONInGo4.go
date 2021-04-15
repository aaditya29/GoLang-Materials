/* Decoding JSON
Decoding JSON is a little bit tricky because we need to coerce some text based daya into a complex data structure.
To decode JSON into a valid data structure like map or struct, we first need to make sure if a JSON is valid.
func Valid(data []byte) bool
We can use json.Valid function to check if JSON is valid. This function returns true if JSON data is valid or false otherwise.
*/

package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//Random JSON Data
	data := []byte(`
	{
		"FirstName": "John",
		"Age": 21,
		"Username": "johndoe91",
		"Grades": null,
		"Languages": [
		  "English",
		  "French"
		]
	}`) //JSON data stored inside data variable

	// check if `data` is valid JSON
	isValid := json.Valid(data) //Using json.Valid(data) function check whether data contains valid JSON
	fmt.Println(isValid)
}
