//Most of the time we want a specific route to handle the static file srving part.
//Creating a static route to serve the content of /tmp directory

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("/Users/Downloads/tmp")) //Creating a file server handler
	// handle `/` route
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "text/html")
		fmt.Fprint(res, "<h1>Golang!</h1>")
	})

	// handle `/static` route
	http.Handle("/static", fs)

	// start HTTP server with `http.DefaultServeMux` handler
	log.Fatal(http.ListenAndServe(":9000", nil))

}
