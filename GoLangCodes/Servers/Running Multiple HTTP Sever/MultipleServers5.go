/* Shutdown method
The server.Shutdown method behaves exactly like server.Close method but with some added safety.
If you want to gracefully close active TCP connections, then Shutdown() is better compared to Close method.

One good advantage of using Shutdown method over Close is that we can perform some cleanup operations after a server shutdown is initiated.
This can be done by registering one or more cleanup functions using RegisterOnShutdown method.
SYNTAX: server.RegisterOnShutdown(f func())
These cleanup functions will be run as separate goroutines and they will be invoked as soon as the shutdown process is started.
Like Close method, Shutdown method can not close upgraded or hijacked connections. Hence this is a good place to close these connections gracefully.

This method needs a base context of the incoming requests. If you haven’t configured the BaseContext field of the Server structure,
you can use the context.Background() since it is the default value for the BaseContext field.
SYNTAX: error := server.Shutdown(ctx context.Context)
*/

package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {

	// create a `WaitGroup` for safe exit
	wg := new(sync.WaitGroup)
	wg.Add(2) // add `2` goroutines to finish

	// create server to run on port the 9000
	server := &http.Server{
		Addr:    ":9000",
		Handler: nil, // use `DefaultServeMux`
	}

	// register a cleanup function
	server.RegisterOnShutdown(func() {
		fmt.Println("RegisterOnShutdown(): completed!") // perform a cleanup
		wg.Done()                                       // WaitGroup done
	})

	// shutdown server after 3 seconds
	time.AfterFunc(3*time.Second, func() {
		err := server.Shutdown(context.Background()) // shutdown `server`
		fmt.Println("Shutdown(): completed!", err)
		wg.Done() // WaitGroup done
	})

	// launch server
	fmt.Println("ListenAndServe():", server.ListenAndServe())

	// listen for `wg` to complete
	wg.Wait()
	fmt.Println("Main(): exit complete!")
}
