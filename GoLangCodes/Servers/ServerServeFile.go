/* ServeFile function:
The http package provides a ServeFile function to serve a file on the disk using its file path.
SYNTAX: func ServeFile(w http.ResponseWriter, r *http.Request, name string)
Here the name argument is the path of the file on the disk.
*/

//Creating a simple HTTP serve to serve some files using ServeFile function

package main

import (
	"log"
	"net/http"
)

// temporary directory location
const TmpDir = "/Users/Downloads/tmp"

func main() {

	// return a `.pdf` file for `/pdf` route
	http.HandleFunc("/pdf", func(res http.ResponseWriter, req *http.Request) {
		http.ServeFile(res, req, TmpDir+"/files/test.pdf") //Using a /pdf route path to get the test.pdf file
	})

	// start HTTP server with `http.DefaultServeMux` handler
	log.Fatal(http.ListenAndServe(":9000", nil))

}
