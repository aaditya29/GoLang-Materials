/* Server type:

To create and configure a custom HTTP server, http package provides Server structure.
This structure has below public fields

type Server struct {
    // TCP address to listen on, ":http" or ":https" if empty
    Addr string
    // http.DefaultServeMux if `nil`
    Handler Handler
    // TLS configuration for HTTPS protocol
    TLSConfig *tls.Config
    // ReadTimeout is the maximum duration for reading the entire
    // request, including the body.
    ReadTimeout time.Duration
    // WriteTimeout is the maximum duration before timing out
    // writes of the response.
    WriteTimeout time.Duration
    // BaseContext optionally specifies a function that returns
    // the base context for incoming requests on this server.
    // If BaseContext is nil, the default is context.Background()
    BaseContext func(net.Listener) context.Context
    // ErrorLog specifies an optional logger for errors accepting
    // connections, unexpected behavior from handlers, and
    // underlying FileSystem errors.
    ErrorLog *log.Logger
}


We are familiar with the Addr and Handler fields. The ReadTimeout and WriteTimeout fields can be important for added safety.
ErrorLog is crucial to log any errors thrown by the server while processing a request.

The Server structure implements basic methods like ListenAndServe, SetKeepAlivesEnabled, Shutdown, etc. to configure and control a server.
*/

//Creating custom srvers and run them concurrently

package main

import (
	"fmt"
	"net/http"
	"sync"
)

func createServer(name string, port int) *http.Server {
	/*
	   Here the createServer function creates a ServerMux object and implements a fresh routing mechanis,.
	   The server object contains the implementation of a brand new HTTP server based on input arguments
	   It then returns a pointer to the newly created Server object since most of the moethods implemented by the Server structure has a pointer reciever.
	*/

	// create `ServerMux`
	mux := http.NewServeMux()

	// create a default route handler
	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Hello: "+name)
	})

	// create new server
	server := http.Server{
		Addr:    fmt.Sprintf(":%v", port), // :{port}
		Handler: mux,
	}

	// return new server (pointer)
	return &server
}

func main() {

	// create a WaitGroup
	wg := new(sync.WaitGroup)

	// add two goroutines to `wg` WaitGroup
	wg.Add(2)

	// goroutine to launch a server on port 9000
	go func() {
		server := createServer("ONE", 9000)
		fmt.Println(server.ListenAndServe()) //Unlike http.ListenAndServe method call the server.ListenAndServe method call does not need any arguments.
		wg.Done()
	}()

	// goroutine to launch a server on port 9001
	go func() {
		server := createServer("TWO", 9001)
		fmt.Println(server.ListenAndServe())
		wg.Done()
	}()

	// wait until WaitGroup is done
	wg.Wait()

}
