/* Decoder
The json/Decoder structure type lets you create a struct that holds a io.Reader object and provides Decode() method to decode JSON from this io.Writer object and write to an object.
SYNTAX: func (dec *Decoder) Decode(v interface{}) error

If all the lines in the io.Reader has been read, then next Decode call returns io.EOF error. But first, we need to create an *Decoder object from a io.Reader using the NewDecoder function.
SYNTAX:func NewDecoder(r io.Reader) *Decoder
Each time the Encode() method is called, JSON is unmarshaled from r by reading a line (trailing with a newline character) and saved to v.
*/
package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Person struct {
	Name string
	Age  int
}

func main() {

	// create a strings reader
	jsonStream := strings.NewReader(`
{"Name":"Ross Geller","Age":28}
{"Name":"Monica Geller","Age":27}
{"Name":"Jack Geller","Age":56}
`)

	// create JSON decoder using `jsonStream`
	decoder := json.NewDecoder(jsonStream)

	// create `Person` structs to hold decoded data
	var ross, monica Person

	// decode JSON from `decoder` one line at a time
	decoder.Decode(&ross)
	decoder.Decode(&monica)

	// see value of the `ross` and `monica`
	fmt.Printf("ross: %#v\n", ross)
	fmt.Printf("monica: %#v\n", monica)
}
