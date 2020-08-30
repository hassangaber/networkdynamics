//Go algorithm which displays a custom file to the server based on a static request & file
//Implementation 3 of GoServe
//Handles http requests, upon '/' it only displays a test string, and upon '/hello.html' it displays
//the html file as written by the author. This can be changed to whatever is in need.
//

package main

import (
	"io"
	"log"
	"net/http"
)

//Function to display string to user upon http request to the "/" directory
func _SendData1(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Test Data to server.\n")
}

//Function which displays the html file upon a request to the "/hello.html" directory
func _CustomFile(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		res,
		`<DOCTYPE html>
  <html>
	<head>
		<title>Example I</title>
	</head>
	<body>
		<h1>Hello!</h2>
	</body>
  </html>`,
	)
}

func main() {

	http.HandleFunc("/", _SendData1)
	http.HandleFunc("/hello.html", _CustomFile)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
