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
