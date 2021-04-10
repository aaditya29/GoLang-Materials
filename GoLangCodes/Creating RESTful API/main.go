/*Creatinf a simple server which can handle HTTP Requests.
First we will create a new file called 'main.go'. In this we will dfine 3 distint functions.
A 'homepage' function that will handle all requests our root URL.
A 'handleRequests' function that will match the URL path hit with a defined function and a 'main' function which will kick off the API.
*/

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article //Declaring a global Articles array that we can then populate in our main function to simulate a database

//Creating '/articles' which which will return all the articles we defined below in main function
//We will create a new REST endpoint which when hit with a 'HTTP GET' request will return all of the articles for our site
func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles) //json.NewEncoder(w).Encode(article) does the job of encoding our articles array into a JSON string and then writing as part of our response.
}

func homePage(w http.ResponseWriter, r *http.Request) { //http. ResponseWriter interface has a Write method which accepts a byte slice and writes the data to the connection as part of an HTTP response.
	fmt.Fprintf(w, "Welcome To The Homepage!")
	fmt.Println("Endpoint Hit: Homepage")
}

func handleRequests() {
	http.HandleFunc("/", homePage) //http.HandleFunc,tells the http package to handle all requests to the web root ("/") with handler.
	//Adding our articles route and mapping it to our returnAllArticles function
	http.HandleFunc("/articles", returnAllArticles)
	log.Fatal(http.ListenAndServe(":10000", nil))
	/*Here log.Fatal is used for an error which occurs during a process which might not be reversible.
	http.ListenAndServe(":10000", nil) function to listen on localhost with port 10000
	*/
}

func main() {

	//updating our 'main' function so that our 'Articles' variable is populated with some dummy data that we can retrieve and modify later on.
	Articles = []Article{
		Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Title: "Hello 2", Desc: "Article Description2", Content: "Article Content2"},
	}
	handleRequests()
}
