package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// vars are placeholders for now. boilerplate code, until I finish them 1-by-1
type contactEdit struct {
	Message  string
	Success  bool
	Feedback string
}

func contactPage(w http.ResponseWriter, r *http.Request) {
	layoutPage := filepath.Join("templates", "layout.gohtml")
	templatePage := filepath.Join("templates", "contact.gohtml")

	// 404 if the template or dir doesn't exist
	info, notFoundErr := os.Stat(templatePage)
	if notFoundErr != nil {
		if os.IsNotExist(notFoundErr) {
			http.NotFound(w, r)
		}
	}
	if info.IsDir() {
		http.NotFound(w, r)
	}
	var templates = template.Must(template.ParseFiles(layoutPage, templatePage))
	var editContact contactEdit
	// Route check
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Println("method:", r.Method)
	// Switch case for GET and POST
	switch r.Method {
	case "GET":
		tmplErr := templates.ExecuteTemplate(w, "layout", nil)
		if tmplErr != nil {
			log.Println(tmplErr.Error())
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		}
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			log.Println(err.Error())
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		editContact = contactEdit{
			Message: r.FormValue("message"),
			Success: true,
			Feedback: "feedback sent",
		}
		// do something with editContact.Message etc. above code is just to test
		tmplErr := templates.ExecuteTemplate(w, "layout", editContact)
		if tmplErr != nil {
			log.Println(tmplErr.Error())
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		}
	default:
		fmt.Fprintf(w, "Only GET and POST methods are supported.")
	}
}
