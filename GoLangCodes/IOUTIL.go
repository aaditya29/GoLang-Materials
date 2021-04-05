/*The io/ioutil package provides simple utility functions to deal with files without having to worry
too much about internal implementation.
Working on directory /Users/Documents/tmp for  demonstration of these functions
*/

//ReadDir function
/* The ioutil.ReadDir function returns the information about files contained in a directory.

SYNTAX: func ReadDir(dirname string) ([]os.FileInfo, error)
This Function takes a string Argument which is the path of the directory
and returns a list of FileInfo objects.
The FileInfo interface defines some important methods to get information about the file.

type FileInfo interface {
    Name() string       // base name of the file
    Size() int64        // length in bytes
    Mode() FileMode     // file mode bits
    ModTime() time.Time // modification time
    IsDir() bool        // abbreviation for Mode().IsDir()
    Sys() interface{}   // underlying data source (can return nil)
}
*/
package main

import (
	"fmt"
	"io/ioutil"
)

//Temporary Directory
const tmpDir = "/Users/Downloads/tmp"

func main() {
	//get files from /tmp directory
	files, _ := ioutil.ReadDir(tmpDir)

	//list information of each file
	for _, file := range files {
		fmt.Printf("Name: %v, Size: %v kb, Mode: %v, IsDir:%v\n",
			file.Name(),
			file.Size()/1000,
			file.Mode(),
			file.IsDir(),
		)
	}
}

/* ReadFile Function

To read a file using a file path, we can use the ioutil.Readfile function.
This function returns the content of the file as an array of bytes.

SYNTAX: func ReadFile(filepath string) ([]byte, error)

*/

package main

import (
    "fmt"
    "io/ioutil"
)

// temporary directory
const tmpDir = "/Users/Downloads/tmp"

func main() {

    // location of `index.html` file
    html, _ := ioutil.ReadFile( "/Users/Uday.Hiwarale/tmp/index.html" )

    // print raw bytes
    fmt.Println( "Raw: \n", html )

    // print string representation
    fmt.Println( "String: \n", string(html) )

}

/* WriteFile function
To simply write some binary data to a file, we can use ioutil.WriteFile function.

SYNTAX: func WriteFile(filepath string, data []byte, perm os.FileMode) error

If a file with the filepath does not exist, it will be created.
If the file already exists, the file’s content will be wiped out (truncated) before writing new content.

*/

package main

import (
    "fmt"
    "io/ioutil"
    "path/filepath"
)

// temporary directory
const tmpDir = "/Users/Uday.Hiwarale/tmp"

func main() {

    // welcome message content
    welcomeData := []byte("Welcome to my website.")

    // get `welcome.txt` file path
    path := filepath.Join( tmpDir, "/welcome.txt" )

    // write content to `welcome.txt` file
    err := ioutil.WriteFile( path, welcomeData, 0777 )

    // log error
    if( err != nil ) {
        fmt.Println(err)
    }

}