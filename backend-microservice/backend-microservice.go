package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// JS and CSS handling/serving
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./templates/css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./templates/js"))))
	// Page serving section
	mux := http.NewServeMux()
	mux.HandleFunc("/", welcomePage)
	mux.HandleFunc("/experience/", experiencePage)
	mux.HandleFunc("/certifications/", certificationPage)
	mux.HandleFunc("/skills/", skillsPage)
	mux.HandleFunc("/recommendations/", recommendationsPage)
	mux.HandleFunc("/accomplishments/", accomplishmentsPage)
	mux.HandleFunc("/projects/", projectsPage)
	mux.HandleFunc("/settings/", settingsPage)
	mux.HandleFunc("/faq/", faqPage)
	mux.HandleFunc("/contact/", contactPage)
	err := http.ListenAndServe(envPort(), mux)
	if err != nil {
		fmt.Println(" Error starting HTTP server...")
		fmt.Println("<=============================>")
		log.Println(err.Error())
	}
}
