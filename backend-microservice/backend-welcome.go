package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type welcomeEdit struct {
	Message  string
	Success  bool
	Feedback string
}

func welcomePage(w http.ResponseWriter, r *http.Request) {
	layoutPage := filepath.Join("templates", "layout.html")
	templatePage := filepath.Join("templates", "welcome.html")

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
	// 500 if it fails to parse the files
	var templates = template.Must(template.ParseFiles(layoutPage, templatePage))
	editMessage := welcomeEdit{
		Message: r.FormValue("message"),
		Success: true,
		Feedback: "feedback sent",
	}
	// editMessage.Success = true
	// editMessage.Feedback = "feedback sent"
	_ = editMessage // make POST to API server and wait for a response
	// execute templates if no post was made
	fmt.Println("method:", r.Method)
	if r.Method != http.MethodPost {
		err := templates.ExecuteTemplate(w, "layout", nil)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		}
	} else {
		tmplErr := templates.ExecuteTemplate(w, "layout", editMessage)
		if tmplErr != nil {
			log.Println(tmplErr.Error())
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		}
	}
}
