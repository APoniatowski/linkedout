package main

import (
	"net/http"
)

const (
	welcome 		= "/"
	experience      = "/experience"
	certifications  = "/certifications"
	skills          = "/skills"
	recommendations = "/recommendations"
	accomplishments = "/accomplishments"
	projects        = "/projects"
	settings        = "/settings"
	faq             = "/faq"
	contact         = "/contact"
)

type dataHandler struct{
	welcomeAPICall welcomeEdit
	contactAPICall contactEdit
	// more to come
}

func (dH *dataHandler)apiPost(r *http.Request) (feedback string) {
	switch r.URL.Path {
	case welcome:
		// http.Post()
		feedback = "POST made"
		return
	case experience:
		// http.Post()
		feedback = "POST made"
		return
	case certifications:
		// http.Post()
		feedback = "POST made"
		return
	case skills:
		// http.Post()
		feedback = "POST made"
		return
	case recommendations:
		// http.Post()
		feedback = "POST made"
		return
	case accomplishments:
		// http.Post()
		feedback = "POST made"
		return
	case projects:
		// http.Post()
		feedback = "POST made"
		return
	case settings:
		// http.Post()
		feedback = "POST made"
		return
	case faq:
		// http.Post()
		feedback = "POST made"
		return
	case contact:
		// http.Post()
		feedback = "POST made"
		return
	default:
		feedback = "POST request made to incorrect URL path."
	}
	/* send data to api and api posts it to DB, and api informs if it was successful
	 and returns the response from api*/
	return
}

func apiGet(r *http.Request) (feedback string) {
	switch r.URL.Path {
	case welcome:

	case experience:

	case certifications:

	case skills:

	case recommendations:

	case accomplishments:

	case projects:

	case settings:

	case faq:

	case contact:

	default:
		feedback = "GET request made to incorrect URL path."
	}
	/* send data to api and api posts it to DB, and api informs if it was successful
	and returns the response from api*/
	return
}

func apiUpdate(r *http.Request) (feedback string) {
	switch r.URL.Path {
	case welcome:

	case experience:

	case certifications:

	case skills:

	case recommendations:

	case accomplishments:

	case projects:

	case settings:

	case faq:

	case contact:

	default:
		feedback = "Update request made to incorrect URL path."
	}
	/* send data to api and api posts it to DB, and api informs if it was successful
	and returns the response from api*/
	return
}

func apiDelete(r *http.Request) (feedback string) {
	switch r.URL.Path {
	case welcome:

	case experience:

	case certifications:

	case skills:

	case recommendations:

	case accomplishments:

	case projects:

	case settings:

	case faq:

	case contact:

	default:
		feedback = "Delete request made to incorrect URL path."
	}
	/* send data to api and api posts it to DB, and api informs if it was successful
	and returns the response from api*/
	return
}
