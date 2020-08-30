//Go algorithm for displaying strings to servers
//Implementation 1 of GoServe
//Handles a http request with a string in the header to display to the user.
//

package main

import (
	"fmt"
	"html"     //displaying html static files on server for later algorithms
	"log"      //logging and processing requests
	"net/http" //http handle functions
)

func main() {

	//http function which returns a string upon a local request.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello", html.EscapeString(r.URL.Path))
	})

	//http function which returns a different string upon a local request. This is also listed
	//as a query.
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "test")
	})

	//error logging for listening issues.
	log.Fatal(http.ListenAndServe(":8081", nil))

}
