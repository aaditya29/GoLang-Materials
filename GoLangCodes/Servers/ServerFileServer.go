/*On the web, we do not just serve processed HTML but files as well, like CSS, JavaScript, PDF, etc.
Most of the time, these files are served as-is from the disk but sometimes, they need to be processed before delivering to the client.

1.http.FileServer():
The http.Fileserver() function provides the functionality to serve the entire file-system directory with indexes.
SYNTAX: func FileServer(root FileSystem) Handler
Here we can see it returns a Handler object which makes it a perfect candidate to be used as handler in ListenAndServe function or the Handle function.
It takes an argument of FileSystem interface type.
The root argument represents the file system directory from which the content will be served.

2. http.Dir type
The Dir type may look like a function but it is an alias of string data type and it implements Open method
defined by the FileSystem interface. We can call the Dir type like a function which is nothing but a type-casting syntax.

The string value passed to the type-casting syntax is a relative or an absolute directory on the native file-system.

var fs FileSystem = http.Dir("/tmp")
*/

//Serving A Directory
//Making a simple Go HTTP Server to serve /Users/Downloads/tmp directory using http.FileServer handler and http.Dir type.
package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("/Users/Downloads/tmp")) //Creating file server handler
	log.Fatal(http.ListenAndServe(":9000", fs))             //Starting a HTTP server with 'fs' as default handler

}
