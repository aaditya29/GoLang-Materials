/* Using Structure Tags
A struct field can hold additional metadata that can be used by other programs to process that field differently.
This metadata is assigned to a field using a string literal (raw string [``] or interpreted string [“”]).
type Data struct {
  FieldOne string `json:"fname" xml:"first-name" gorm:"size:255"`
  FieldTwo string `json:"lname" xml:"last-name" gorm:"size:255"`
}

In our case, Marshal function uses the tag of a struct field to obtain additional encoding/encoding information from the field.
For JSON encoding, we need to use json:"options" tag value. Here, the options are comma-separated string values.

The first option value is the name of the field that should appear in the JSON.
The other option values can be omitempty to discard a field if its value is empty or string to convert the field’s value to a string.
*/

package main

import (
	"encoding/json"
	"fmt"
)

// Profile declares `Profile` structure
type Profile struct {
	Username  string `json:"uname"`
	Followers int    `json:"followers,omitempty,string"`
}

// Student declares `Student` structure
type Student struct {
	FirstName string  `json:"fname"`           // `fname` as field name
	LastName  string  `json:"lname,omitempty"` // discard if value is empty
	Email     string  `json:"-"`               // always discard
	Age       int     `json:"-,"`              // `-` as field name
	IsMale    bool    `json:",string"`         // keep original field name, coerce to a string
	Profile   Profile `json:""`                // no effect
}

func main() {

	// define `john` struct (pointer)
	john := &Student{
		FirstName: "John",
		LastName:  "", // empty
		Age:       21,
		Email:     "john@doe.com",
		Profile: Profile{
			Username:  "johndoe91",
			Followers: 1975,
		},
	}

	// encode `john` as JSON
	johnJSON, _ := json.MarshalIndent(john, "", "  ")

	// print JSON string
	fmt.Println(string(johnJSON))
}
