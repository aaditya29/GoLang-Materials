/* HTTP Server:
An HTTP server is a program that listens for HTTP requests on a port of the machine
with a static or dynamic IP address. The HTTP protocol is an internet protocol
made up of TCP and IP protocol.

When a browser (client) makes an HTTP request to a machine with a given address and port,
it sends the data in packets to the machine. These packets are assembled by the HTTP server and
the server may choose to respond with some data requested by the client over the same request connection.

While creating an HTTP server, we need to make sure that we are binding the server on the correct port.
If the port isn’t correct, then all the requests made by the client on the port won’t reach the server because the server is listening on some other port.
*/

package main

import (
	"net/http"
)

type HttpHandler struct{} //Create a handler struct

func (h HttpHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) { //Implementing 'ServerHTTP' method on 'HttpHandler' struct
	/*
				Anatomy of ServeHTTP method
				When the HTTP server is started by invoking ListenAndServe(addr, handler) function, any incoming HTTP request will trigger
				the ServeHTTP method of the handler argument.
				The main job of ServeHTTP method is to respond to the request with some data. This method has the following signature:
				ServeHTTP(res http.ResponseWriter, req *http.Request)
				The res argument contains the response object of interface type whereas ResponseWriter has the following description:
				type ResponseWriter interface {
		  	  	Header() Header
		    	Write([]byte) (int, error)
		    	WriteHeader(statusCode int)
				}
	*/

	data := []byte("Hello World!") //Creating response binary data with slice of bytes

	res.Write(data) //Writing 'data' to response

}

func main() {
	handler := HttpHandler{} //Creating a new handler

	http.ListenAndServe(":9000", handler)
	/*
		The function ListernAndServe starts and HTTP server and locks the process.
		It starts to listen to incoming HTTP requests and serve requests using Serve function under the hood.
		ListenAndServe internally creates a tcp listener on address addr using net.Listen function which returns a
		net.Listener and uses it with http.Serve function to listen to incoming connections using handler.

		Since this function locks the provess, any code below this function call will not be executed
	*/

}
