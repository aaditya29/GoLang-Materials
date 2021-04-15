/*
Since a JSON contains string keys and values of supported data types, a map of type map[string]interface{} is a suitable candidate for storing JSON data.
We can pass a pointer to nil or non-nil pointer of the map to the Unmarshal function and all JSON field values will be populated inside the map.
*/

//Working With Maps

package main

import (
	"encoding/json"
	"fmt"
)

// Student declares `Student` map
type Student map[string]interface{}

func main() {

	// some JSON data
	data := []byte(`
	{
		"id": 123,
		"fname": "John",
		"height": 1.75,
		"male": true,
		"languages": null,
		"subjects": [ "Math", "Science" ],
		"profile": {
			"uname": "johndoe91",
			"f_count": 1975
		}
	}`)

	// create a data container
	var john Student

	// unmarshal `data`
	fmt.Printf("Error: %v\n", json.Unmarshal(data, &john))

	// print `john` map
	fmt.Printf("%#v\n\n", john)

	// iterate through keys and values
	i := 1
	for k, v := range john {
		fmt.Printf("%d: key (`%T`)`%v`, value (`%T`)`%#v`\n", i, k, k, v, v)
		i++
	}
}

/*
There are certain rules Unmarshal functions follows to store the JSON values in a map.
1. A JSON string value is stored as string.
2. A JSON number value(int or float) is stored as float64.
3. A JSON boolean value is stored as bool.
4. A JSON null value is stored as nil value.
5. A JSON array value is stored as a slice of type []interface{}.
6. A JSON object value is stored as a map of type map[string]interface{}.
*/
