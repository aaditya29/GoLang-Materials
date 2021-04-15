/* JSON Unmarshalling
Like we used the json.Marshal function to encode JSON data from a data structure, we have to use json.Unmarshal function
to decode JSON data into a data structure like map or struct.
SYNTAX: func Unmarshal(data []byte, v interface{}) error
The Unmarshal function takes the JSON data as the first argument and the container v that will hold the data as the second argument.
The v argument is either a pointer to a valid data structure or to an interface.
*/

package main

import (
	"encoding/json"
	"fmt"
)

//Declaring 'Student' Structure
type Student struct {
	FirstName, lastName string
	Email               string
	Age                 int
	HeightInMeters      float64
}

func main() {
	//Random JSON Data
	data := []byte(`
	{
		"FirstName": "John",
		"lastName": "Doe",
		"Age": 21,
		"HeightInMeters": 175,
		"Username": "johndoe91"
	}`)

	//Creating a data container
	var john Student //Declaring an empty struct john of type Student
	//As things go with struct, an empty struct will have all its fields set to their respective zero-values (like 0 for integers and “” for strings).

	//Unmarshalling 'data'
	fmt.Printf("Error: %v\n", json.Unmarshal(data, &john))

	//Printing 'John' struct
	fmt.Printf("%#v\n", john)
}
