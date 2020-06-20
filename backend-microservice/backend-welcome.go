package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type welcomeEdit struct {
	Message string
}

func welcomePage(w http.ResponseWriter, r *http.Request) {
	layoutPage := filepath.Join("templates", "layout.html")
	templatePage := filepath.Join("templates", "welcome.html")

	// 404 if the template or dir doesn't exist
	info, notFoundErr := os.Stat(templatePage)
	if notFoundErr != nil {
		if os.IsNotExist(notFoundErr) {
			http.NotFound(w, r)
			return
		}
	}
	if info.IsDir() {
		http.NotFound(w, r)
		return
	}

	templates, parseErr := template.ParseFiles(layoutPage, templatePage)
	if parseErr != nil {
		log.Println(parseErr.Error())
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
	}
	if r.Method != http.MethodPost {
		err := templates.Execute(w, nil)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		}
		return
	}
	editMessage := welcomeEdit{
		Message: r.FormValue("message"),
	}

	_ = editMessage // make POST to API server and wait for a response

	tmplErr := templates.Execute(w, struct{ Success bool }{true})
	if tmplErr != nil {
		log.Println(tmplErr.Error())
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
	}
}
