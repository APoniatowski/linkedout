package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/welcome", dbWelcome)
	err := http.ListenAndServe(envPort(), nil)
	if err != nil {
		fmt.Println(" Error starting HTTP server...")
		fmt.Println("<=============================>")
		fmt.Println(err)
	}
}
