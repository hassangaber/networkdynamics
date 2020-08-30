package main

import (
	"fmt"
	"net/http"
)

func main() {

	fmt.Println("Go has been successfully installed and ready to be used for Offset tools.")
	fmt.Println("However, we must do some tests first.\n Just an Internet Connection test.")

	_InternetConnection()
}
func _InternetConnection() (Connection bool) {
	_, err := http.Get("http://clients3.google.com/generate_204")
	if err != nil {
		fmt.Println("No Internet Connection.")
		return false
	} else {
		fmt.Println("Positive Internet Connection")
		return true
	}
}
