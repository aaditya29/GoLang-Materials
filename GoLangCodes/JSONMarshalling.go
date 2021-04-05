/*JSON is a simple data interchange format.
It is most commonly used for communication between web backends and JS porgorams running in the browser.

With the json package it's a snap to read and write JSON data from GO Programs.
*/
package main
//Marshal function is used to encode the JSON Data.
func Marshal(v interface{}) ([]byte, error)

//GO Data Structure Messages
type Message struct {
	Name string
	Body string 
	Time int64
}

//Instance of message
m := Message{"Alice", "Hello", 1294706395881547000}

//We can marshal a JSON-encoded version of m using json.Marshal
b, err := json.Marshal(m)

b == []byte(`{"Name":"Alice","Body":"Hello","Time":1294706395881547000}`)