package main

import (
	"net/http"
)

/*
	placeholder db query. this will send a soon to be defined request to the DB server.
	mongoDB will be used, to avoid json marshalling and unmarshalling.
*/
func Welcome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/welcome" {
		http.NotFound(w, r)
		return
	}
	switch r.Method {
	case "GET":
		// should check DB if anything exists, if it exists, send it
	case "POST":
		// if something is posted, it overwrite the welcome entry. Wasted resources trying to update. overwriting is cheaper
	case "PUT":
		// send to POST. Overwrite is the cheaper option
	case "DELETE":
		// I doubt I need to explain this one, to anyone...
	default:
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
	}
}
