package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// Pages that will be served
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

	tmplErr := templates.ExecuteTemplate(w, "layout", nil)
	if tmplErr != nil {
		log.Println(tmplErr.Error())
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
	}
}

func experiencePage(w http.ResponseWriter, r *http.Request) {
	layoutPage := filepath.Join("templates", "layout.html")
	templatePage := filepath.Join("templates", "experience.html")

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

func certificationPage(w http.ResponseWriter, r *http.Request) {
	layoutPage := filepath.Join("templates", "layout.html")
	templatePage := filepath.Join("templates", "certifications.html")

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

func skillsPage(w http.ResponseWriter, r *http.Request) {
	layoutPage := filepath.Join("templates", "layout.html")
	templatePage := filepath.Join("templates", "skills.html")

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

func recommendationsPage(w http.ResponseWriter, r *http.Request) {
	layoutPage := filepath.Join("templates", "layout.html")
	templatePage := filepath.Join("templates", "recommendations.html")

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

func accomplishmentsPage(w http.ResponseWriter, r *http.Request) {
	layoutPage := filepath.Join("templates", "layout.html")
	templatePage := filepath.Join("templates", "accomplishments.html")

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

func contactPage(w http.ResponseWriter, r *http.Request) {
	layoutPage := filepath.Join("templates", "layout.html")
	templatePage := filepath.Join("templates", "contact.html")

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
