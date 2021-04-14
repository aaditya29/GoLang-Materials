/* As we know once we run http.ListenAndServe() function, our process gets blocked and we can't run anything below this time.

When we call http.ListenAndServer() with appropriate arguments, it starts an HTTP server and blocks the current goroutine.
If we run this function inside the main function, it will block the main goroutine (started by the main function).
Once you call ListenAndServe function, you do not have another chance to spawn a new server, unless you run a goroutine that launches another server.

Now we will see how to achieve it in following program
*/

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	//Creating a default route handler
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Hello: "+req.Host)
	})

	//Creating a goroutine
	go func() {
		log.Fatal(http.ListenAndServe(":9000", nil)) //Spawning an HTTP server in another goroutine
	}()

	log.Fatal(http.ListenAndServe("9001", nil)) //Spawning HTTP server in 'main' goroutine
}
