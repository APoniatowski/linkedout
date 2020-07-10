package main

import (
	"fmt"
	"net/http"

)

func main() {

	linkedoutHandlers := newLinkedoutHandlers()
	http.HandleFunc("/welcome", linkedoutHandlers.Welcome)
	err := http.ListenAndServe(envPort(), nil)
	if err != nil {
		fmt.Println(" Error starting HTTP server...")
		fmt.Println("<=============================>")
		fmt.Println(err)
	}
}
