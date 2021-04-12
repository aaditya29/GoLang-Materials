/* We can pass nil as the value of handler to the ListenAndServe function.
When we pass nil, Go will internally use the http.DefaultServeMux which is a global ServeMux instance.
This will make our life a little easier as we don’t have to create ServeMux manually.

We can add routes and handler functions to this instance using similar methods but
using the functions exposed by http package

The HandleFunc function provided by the http package does the exact same thing that the mux.HandleFunc does.
It adds a handler function to the http.DefaultServeMux to handle HTTP requests of a specific route path.
*/

//Passing a nil handler
package main

import (
	"fmt"
	"net/http"
)

func main() {

	// handle `/` route to `http.DefaultServeMux`
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Hello World!")
	})

	// handle `/hello/golang` route to `http.DefaultServeMux`
	http.HandleFunc("/hello/golang", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Hello Golang!")
	})

	// listen and serve using `http.DefaultServeMux`
	http.ListenAndServe(":9000", nil)

}
