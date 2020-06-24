package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Page serving section
	mux := http.NewServeMux()
	mux.HandleFunc("/", welcomePage)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./templates/static"))))
	mux.HandleFunc("/experience", experiencePage)
	mux.HandleFunc("/certifications", certificationPage)
	mux.HandleFunc("/skills", skillsPage)
	mux.HandleFunc("/recommendations", recommendationsPage)
	mux.HandleFunc("/accomplishments", accomplishmentsPage)
	mux.HandleFunc("/projects", projectsPage)
	mux.HandleFunc("/settings", settingsPage)
	mux.HandleFunc("/faq", faqPage)
	mux.HandleFunc("/contact", contactPage)
	err := http.ListenAndServe(envPort(), mux)
	if err != nil {
		fmt.Println(" Error starting HTTP server...")
		fmt.Println("<=============================>")
		log.Println(err.Error())
	}
}
