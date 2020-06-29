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
	IsPOST  bool
	Feedback string
}

func welcomePage(w http.ResponseWriter, r *http.Request) {
	layoutPage := filepath.Join("templates", "layout.gohtml")
	templatePage := filepath.Join("templates", "welcome.gohtml")

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
	var editMessage dataHandler
	// Route check
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
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
		editMessage.welcomeAPICall = welcomeEdit{
			Message: r.FormValue("message"),
			IsPOST: true,
		}
		editMessage.welcomeAPICall.Feedback = editMessage.apiPost(r)
		// do something with editMessage.Message etc. above code is just to test
		tmplErr := templates.ExecuteTemplate(w, "layout", editMessage.welcomeAPICall)
		if tmplErr != nil {
			log.Println(tmplErr.Error())
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		}
	default:
		fmt.Fprintf(w, "Only valid methods are supported.")
	}
}
