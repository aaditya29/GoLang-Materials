/*
Write a program which prompts the user to first enter a name, and then enter an address.
Your program should create a map and add the name and address to the map using the keys “name” and “address”, respectively.
Your program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object.
*/

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var m map[string]string
	m = make(map[string]string)
	fmt.Println("Enter your name:")
	br := bufio.NewReader(os.Stdin)
	bname, _, _ := br.ReadLine()
	name := string(bname)
	m["name"] = name
	fmt.Println("Input your address: ")
	br1 := bufio.NewReader(os.Stdin)
	baddress, _, _ := br1.ReadLine()
	address := string(baddress)
	m["address"] = address
	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println("Encoding Failed.")
	} else {
		fmt.Println("Encoded data: ")
		fmt.Println(b)
		fmt.Println("Decoded data: ")
		fmt.Println(string(b))
	}

}
