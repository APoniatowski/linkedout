package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", welcomePage)
	http.HandleFunc("/experience/", experiencePage)
	http.HandleFunc("/certificates/", certificationPage)
	http.HandleFunc("/skills/", skillsPage)
	http.HandleFunc("/recommendations/", recommendationsPage)
	http.HandleFunc("/accomplishments/", accomplishmentsPage)
	http.HandleFunc("/faq/", faqPage)
	http.HandleFunc("/contact/", contactPage)
	err := http.ListenAndServe(envPort(), nil)
	if err != nil {
		fmt.Println(" Error starting HTTP server...")
		fmt.Println("<=============================>")
		fmt.Println(err)
	}
}

// Pages that will be served
func welcomePage(w http.ResponseWriter, r *http.Request) {

	type Testslice struct {
		Title    string
		SubTitle bool
	}

	type PageData struct {
		PageTitle string
		LOData    []Testslice
	}

	w.WriteHeader(http.StatusOK)
	welcomeTemplate := template.Must(template.ParseFiles("./templates/welcome.html"))

	data := PageData{
		PageTitle: "Welcome!",
		LOData: []Testslice{
			{Title: "Data 1", SubTitle: false},
			{Title: "Data 2", SubTitle: true},
			{Title: "Data 3", SubTitle: true},
		},
	}
	err := welcomeTemplate.Execute(w, data)
	if err != nil {
		fmt.Println("Error executing template")
	}
}

func experiencePage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Experience Page")
}

func certificationPage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Certifications Page")
}

func skillsPage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Skills Page")
}

func recommendationsPage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Recommendations Page")
}

func accomplishmentsPage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Accomplishments Page")
}

func faqPage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "FAQ Page")
}

func contactPage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Contact me Page")
}

// Port environment variable for the frontend. This can be set to something else in the container's ENV, as PORT. Default is 80 (HTTP)
func envPort() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "80"
	}
	return ":" + port
}
