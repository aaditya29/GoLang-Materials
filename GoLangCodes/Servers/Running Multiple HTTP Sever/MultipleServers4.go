/* Close method in a server
Once we no longer need a server, we can close it without terminating the main program. This is achieved using the server.Close method.
This method forcefully shutdown a server without gracefully closing the active TCP connection.
SYNTAX: error := server.Close()

Therefore this method can only be used when a server needs to be shut down immediately without caring about active TCP connections.
*/

package main

import (
	"fmt"
	"net/http"
	"time"
)

func maine() {
	//Creating channels for safe exit
	exitSignal := make(chan interface{})

	//Creating server to run on port the 9000
	server := &http.Server{
		Addr:    "9000",
		Handler: nil, //Using 'DefaultServeMux'
	}

	//Closing server after 3 seconds
	time.AfterFunc(3*time.Second, func() {
		fmt.Println("Close(): completed!", server.Close()) //Closing 'Server'
	})

	//Launching server
	err := server.ListenAndServe()
	fmt.Println("ListenAndServe():", err)

	//Listening to 'exitSignal' channel
	<-exitSignal
	fmt.Println("Main(): exit complete!")

	//The only drawback of Close method is that it can not close hijacked connections like WebSockets.
}
