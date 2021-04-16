/* Encoder
The json/Encoder structure type lets you create a struct that holds a io.Writer object and provides Encode() method to encode JSON from an object and write to this io.Writer object.
SYNTAX: func (enc *Encoder) Encode(v interface{}) error

But first, we need to create an *Encoder object from a io.Writer using the NewEncoder function.
SYNTAX: func NewEncoder(w io.Writer) *Encoder
Each time the Encode() method is called, JSON is marshaled from v and appended to the w with a trailing newline.
*/
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {

	// create a buffer to hold JSON data
	buf := new(bytes.Buffer)
	// create JSON encoder for `buf`
	bufEncoder := json.NewEncoder(buf)

	// encode JSON from `Person` structs
	bufEncoder.Encode(Person{"Ross Geller", 28})
	bufEncoder.Encode(Person{"Monica Geller", 27})
	bufEncoder.Encode(Person{"Jack Geller", 56})

	// print contents of the `buf`
	fmt.Println(buf) // calls `buf.String()` method

}
