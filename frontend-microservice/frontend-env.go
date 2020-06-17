package main

import "os"

// Port environment variable for the frontend. This can be set to something else in the container's ENV, as PORT. Default is 80 (HTTP)
func envPort() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "80"
	}
	return ":" + port
}
