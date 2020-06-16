package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", apiserver)
	http.HandleFunc("/api/dbquery", dbquery)
	http.ListenAndServe(envPort(), nil)

	// will implement error checking when starting up the listener
	// if err != nil {
	// 	fmt.Println("Error starting HTTP server...")
	// }
}

func envPort() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}

func apiserver(w http.ResponseWriter, r http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Up and running")
}

func dbquery(w http.ResponseWriter, r http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Up and running")
}
