/* Using Unmarshler and TextUnmarshaler
A struct field can take responsibility for unmarshaling the JSON data on its own. In such a case, the field value must implement the json.Unmarshaler interface which provides the declaration of UnmarshalJSON method.
type Unmarshaler interface {
    UnmarshalJSON([]byte) error
}
This method is used to delegate the responsibility of unmarshalling a field back to the field itself. If Unmarshal finds a field of type Unmarshaler, it will call the UnmarshalJSON function with JSON data of that field
(even for null) and it becomes the responsibility of that field to initialize/assign a value.
*/

package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Profile declares `Profile` structure
type Profile struct {
	Username  string
	Followers string
}

// UnmarshalJSON - implement Unmarshaler interface
func (p *Profile) UnmarshalJSON(data []byte) error {

	// unmarshal JSON
	var container map[string]interface{}
	_ = json.Unmarshal(data, &container)
	fmt.Printf("container: %T / %#v\n\n", container, container)

	// extract interface values
	iuserName, _ := container["Username"]
	ifollowers, _ := container["f_count"]
	fmt.Printf("iuserName: %T/%#v\n", iuserName, iuserName)
	fmt.Printf("ifollowers: %T/%#v\n\n", ifollowers, ifollowers)

	// extract concrete values
	userName, _ := iuserName.(string)    // get `string` value
	followers, _ := ifollowers.(float64) // get `float64` value
	fmt.Printf("userName: %T/%#v\n", userName, userName)
	fmt.Printf("followers: %T/%#v\n\n", followers, followers)

	// assign values
	p.Username = strings.ToUpper(userName)
	p.Followers = fmt.Sprintf("%.2fk", followers/1000)

	return nil
}

// Student declares `Student` structure
type Student struct {
	FirstName string
	Profile   Profile
}

func main() {

	// some JSON data
	data := []byte(`
	{
		"FirstName": "John",
		"Profile": {
			"Username": "johndoe91",
			"f_count": 1975
		}
	}`)

	// create a data container
	var john Student

	// unmarshal `data`
	fmt.Printf("Error: %v\n", json.Unmarshal(data, &john))

	// print `john` struct
	fmt.Printf("%#v\n", john)
}
