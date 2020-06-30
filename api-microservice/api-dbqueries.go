package main

import (
	"fmt"
	"net/http"
)

/*
	placeholder db query. this will send a soon to be defined request to the DB server.
	mongoDB will be used, to avoid json marshalling and unmarshalling.
*/
func dbWelcome(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Up and running")
}
