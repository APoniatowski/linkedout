package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/api/dbquery", dbquery)
	err := http.ListenAndServe(envPort(), nil)
	if err != nil {
		fmt.Println(" Error starting HTTP server...")
		fmt.Println("<=============================>")
		fmt.Println(err)
	}
}
