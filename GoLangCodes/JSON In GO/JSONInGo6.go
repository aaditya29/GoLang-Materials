/*
If a JSON contains complex data like object or array, then a structure must declare the fields of appropriate types in order to unmarshal the JSON without an error.
*/

//Handling Complex Data

package main

import (
	"encoding/json"
	"fmt"
)

// Profile declares `Profile` structure
type Profile struct {
	Username  string
	Followers int
}

// Student declares `Student` structure
type Student struct {
	FirstName, lastName string
	HeightInMeters      float64
	IsMale              bool
	Languages           [2]string
	Subjects            []string
	Grades              map[string]string
	Profile             Profile
}

func main() {

	// some JSON data
	data := []byte(`
	{
		"FirstName": "John",
		"HeightInMeters": 1.75,
		"IsMale": null,
		"Languages": [ "English", "Spanish", "German" ],
		"Subjects": [ "Math", "Science" ],
		"Grades": { "Math": "A" },
		"Profile": {
			"Username": "johndoe91",
			"Followers": 1975
		}
	}`)

	// create a data container
	var john Student = Student{
		IsMale:   true,
		Subjects: []string{"Art"},
		Grades:   map[string]string{"Science": "A+"},
	}

	// unmarshal `data`
	fmt.Printf("Error: %v\n", json.Unmarshal(data, &john))

	// print `john` struct
	fmt.Printf("%#v\n", john)
}
