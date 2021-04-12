/* Returning A Better Response

Till now our server has become capable of handling routed requests but we haven’t improved
our response structure significantly. A normal HTTP response contains headers to inform the browser about the response.

Following are the methods through which we can add some header values in the response and change the status code of the network response as well:

1) The ResponseWriter.Header() returns the object of Header type.
This object contains the map of header keys and values. We can use Add, Set and Del methods to manipulate the response headers.
We can use the Get method to read existing response headers.
2) The ResponseWriter.WriteHeader(int) method writes the HTTP status code header to the response.
However, this method is called implicitly by Go when first ResponseWriter.Write() call is made with status code 200 (http.StatusOK constant value).
We should pass valid HTTP status integer code or use constants provided by the http package.
*/

/* In this program we will add Content-Type header and reset the Date header from the response. This time, we will send a JSON string as a response.
To notify the browser that the incoming response is of the JSON type, we will use application/json as the value of Content-Type header.
*/

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	// handle `/` route to `http.DefaultServeMux`
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {

		// get response headers
		header := res.Header()

		// set content type header
		header.Set("Content-Type", "application/json")

		// reset date header (inline call)
		res.Header().Set("Date", "01/01/2020")
		/*
			Here we accessed the headers from the response using res.Header() method.
			We can manipulate the response headers from header object or we can make an inline call like in line 40
		*/

		// set status header
		res.WriteHeader(http.StatusBadRequest) // http.StatusBadRequest == 400
		// Here we used res.WriteHeader() method to add the HTTP status code header to the response.
		// respond with a JSON string
		fmt.Fprint(res, `{"status":"FAILURE"}`)
	})

	// listen and serve using `http.DefaultServeMux`
	log.Fatal(http.ListenAndServe(":9000", nil))

}
