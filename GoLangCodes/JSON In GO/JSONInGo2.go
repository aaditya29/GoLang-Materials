/* Data Types Handling
JSON supports primarily 6 data types viz. string, number, boolean, null, array and object.
Apart from these built-in data types, other data types like func, chan, complex64/128 can not be converted to JSON values.
Trying to marshal an object which contains these data types will throw an json.UnsupportedTypeError error except when the field is unconditionally omitted using a tag.
*/

//Encoding Abstract Data Types
package main

import (
	"encoding/json"
	"fmt"
)

//Declaring structure 'Profile'
type Profile struct {
	Username  string
	followers int
	Grades    map[string]string
}

//Declaring structure 'Student
type Student struct {
	FirstName, lastName string
	Age                 int
	Profile             Profile
	Languages           []string
}

func main() {

	var john Student

	// define `john` struct
	john = Student{
		FirstName: "John",
		lastName:  "Doe",
		Age:       21,
		Profile: Profile{
			Username:  "johndoe91",
			followers: 1975,
			Grades:    map[string]string{"Math": "A", "Science": "A+"},
		},
		Languages: []string{"English", "French"},
	}

	// encode `john` as JSON
	johnJSON, err := json.MarshalIndent(john, "", "  ")

	// print JSON string
	fmt.Println(string(johnJSON), err)
}
