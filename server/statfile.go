//Go algorithm to add static files in directory to a server
//Implementation 2 of GoServe
//Adds all static file pings requested from the same directory
//To add a static file to your server, run this file and go to http://localhost:8080/[filename]
//

package main

import (
	"fmt"
	"log"      //logging and processing requests
	"net/http" //http handle functions
)

func main() {

	//Function to process URL request and handle further thus serving the static file in the directory
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})

	//Function which writes string to server for echo purposes
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Tstsrvr")
	})

	//Return nil error upon fatal request to the 8080 local host
	log.Fatal(http.ListenAndServe(":8080", nil))

}
