package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"time"
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

func logError(e error) {
	log.Println(e.Error())
}

type dataHandler struct{
	welcomeAPICall welcomeEdit
	contactAPICall contactEdit
	// more to come
}

func (dH *dataHandler)apiPost(r *http.Request) (feedback string) {
	switch r.URL.Path {
	case welcome:
		// convert data to byte
		welcomePost := []byte(dH.welcomeAPICall.Message)

		// send the converted request
		request, postErr := http.NewRequest("POST", "http://linkedout-api/welcome", bytes.NewBuffer(welcomePost))
		logError(postErr)
		request.Header.Set("Content-type","")

		// set a timeout period
		timeout := 2 * time.Second
		client := http.Client{
			Timeout: timeout,
		}
		// handle the response
		response, resErr := client.Do(request)
		logError(resErr)
		defer response.Body.Close()

		// api will handle json.marshalling and will return json from api
		// it will lessen the burden of marshalling on the frontend, and let the api handle marshalling

		// read the response
		body, bErr := ioutil.ReadAll(response.Body)
		logError(bErr)

		// return the feedback as string
		feedback = string(body)
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
