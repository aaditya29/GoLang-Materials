/* Encoding JSON
To encode a JSON from a suitable data structure, we use json.Marshal function provided by the JSON package
SYNTAX: func Marshal(v interface{}) ([]byte, error)

We can use a struct or map object as the argument v to the Marshal function (to encode JSON data from).
This function returns a slice of bytes which is nothing but the UTF-8 encoded JSON data and an error if
the object v can not be encoded as JSON string due to some reasons.
*/

//Creating a simple struch and encode JSON from it:
package main

import (
	"encoding/json"
	"fmt"
)

//Declaring structure named 'Student'
type Student struct {
	FirstName, lastName string
	Email               string
	Age                 int
	HeightInMeters      float64
	IsMale              bool
}

func main() {
	//defining struct 'john'
	john := Student{
		FirstName:      "John",
		lastName:       "Doe",
		Age:            21,
		HeightInMeters: 1.75,
		IsMale:         true,
	}
	//encoding 'john' as JSON
	johnJson, _ := json.Marshal(john) //Passing john struct to the json.Marshal function and it should return JSON data in slice of bytes

	//Printing JSON string
	//Since we want to print johnJSON as a string, we have used type conversion in the Println function
	fmt.Println(string(johnJson))
}
