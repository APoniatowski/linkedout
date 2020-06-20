package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func faqPage(w http.ResponseWriter, r *http.Request) {
	layoutPage := filepath.Join("templates", "layout.html")
	templatePage := filepath.Join("templates", "faq.html")

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

	tmplErr := templates.ExecuteTemplate(w, "layout", nil)
	if tmplErr != nil {
		log.Println(tmplErr.Error())
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
	}
}
