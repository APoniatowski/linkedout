package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/welcome", Welcome)
	err := http.ListenAndServe(envPort(), nil)
	if err != nil {
		fmt.Println(" Error starting HTTP server...")
		fmt.Println("<=============================>")
		fmt.Println(err)
	}
}
