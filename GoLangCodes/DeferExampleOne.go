package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	//Opening Resource
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close() //Closing of resource
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)

}
