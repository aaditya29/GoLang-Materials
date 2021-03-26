/*Write a program which reads information from a file and represents it in a slice of structs. Assume that there is a text file which contains a series of names. Each line of the text file has a first name and a last name, in that order, separated by a single space on the line.

Your program will define a name struct which has two fields, fname for the first name, and lname for the last name. Each field will be a string of size 20 (characters).

Your program should prompt the user for the name of the text file. Your program will successively read each line of the text file and create a struct which contains the first and last names found in the file.
Each struct created will be added to a slice, and after all lines have been read from the file, your program will have a slice containing one struct for each line in the file.
After reading all lines from the file, your program should iterate through your slice of structs and print the first and last names found in each struct.
*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func main() {
	path := ""
	fmt.Println("Please enter the full path name of the file: \n")
	fmt.Scanln(&path)
	var name []Name
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	for {
		line, _, err := rd.ReadLine()

		if err != nil || io.EOF == err {
			break
		}

		tname := strings.Split(string(line), " ")
		tpname := Name{
			fixLongName(tname[0]),
			fixLongName(tname[1]),
		}
		name = append(name, tpname)

	}
	for i := 0; i < len(name); i++ {
		fmt.Println(i + 1)
		fmt.Println("First Name: " + name[i].fname)
		fmt.Println("Last Name: " + name[i].lname)
	}
}

//cutting the string if name is extra long
func fixLongName(buffer string) string {
	if len(buffer) > 20 {
		return string(buffer[0:20])
	} else {
		return buffer
	}
}
