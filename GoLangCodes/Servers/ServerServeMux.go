/* Understanding ServeMux
The ServeMux is a built-in struct type exported from the http package, that acts as HTTP request multiplexer.
We can create an instance of ServeMux and pass it as the handler to the ListenAndServe() call.
This can be possible because ServeMux implements the ServeHTTP method which will make it implement the Handler interface.

Why We Need ServeMux?
The problem with ServeHTTP method of HttpHandler is that it responds to all the requests.
If you are a web developer, then you always make precise requests to get a specific response,
for example, /index.html to get HTML document and /styles.css to get CSS stylesheet.
The way ServeMux does is to check if the incoming request matches the route it has, as close as possible,
and then execute the function it has associated with that route to handle the response.
*/

package main

import (
	"fmt"
	"net/http"
)

func main() {
	/*
		To get a ServeMux instance, we need to call http.NewServeMux function.
		This function returns an instance of ServeMux with the loaded capability to handle requests based on routes.
	*/
	mux := http.NewServeMux() //Creating a new 'ServeMux'

	/*Now we are creating a new ServeMux object to handle all incoming HTTP requests.
	Then we added two handler functions for / and /hello/golang routes to respond to the HTTP requests received by the mux.
	*/

	//Handling '/' route
	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Hello World!")
	})

	//Handling '/hello/holang' route
	mux.HandleFunc("/hello/holang", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Hello Golang!")
	})

	//Listening and Serving using 'ServeMux'
	http.ListenAndServe(":9000", mux)
}
