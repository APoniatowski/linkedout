package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// JS and CSS handling/serving
	cssStatic := http.FileServer(http.Dir("./templates/css"))
	jsStatic := http.FileServer(http.Dir("./templates/js"))
	http.Handle("/css/", http.StripPrefix("/css/", cssStatic))
	http.Handle("/js/", http.StripPrefix("/js/", jsStatic))

	// Page serving section
	http.HandleFunc("/", welcomePage)
	http.HandleFunc("/experience/", experiencePage)
	http.HandleFunc("/certifications/", certificationPage)
	http.HandleFunc("/skills/", skillsPage)
	http.HandleFunc("/recommendations/", recommendationsPage)
	http.HandleFunc("/accomplishments/", accomplishmentsPage)
	http.HandleFunc("/faq/", faqPage)
	http.HandleFunc("/contact/", contactPage)
	err := http.ListenAndServe(envPort(), nil)
	if err != nil {
		fmt.Println(" Error starting HTTP server...")
		fmt.Println("<=============================>")
		log.Println(err.Error())
	}
}
