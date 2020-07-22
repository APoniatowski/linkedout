package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"sync"
)

const (
	sendSuccess = "success"
	sendError = "error"
)



}

func newLinkedoutHandlers() *linkedoutHandlers {
	return &linkedoutHandlers{
		welcomeHandlers: WelcomeJSON{},
		contactHandlers: ContactJSON{},
	}
}

type linkedoutHandlers struct {
	sync.Mutex
	welcomeHandlers WelcomeJSON
	contactHandlers ContactJSON
}

func (h *linkedoutHandlers) get(w http.ResponseWriter, r *http.Request) {
	// TODO  receive request

	getMessage := h.getMongoDB(strings.TrimLeft(r.URL.Path,"/"))
	h.Lock()
	jsonBytes, marshalErr := json.Marshal(getMessage)
	if marshalErr != nil {
		log.Println(marshalErr.Error())
		w.WriteHeader(http.StatusInternalServerError)
		_ , errErr := w.Write([]byte(marshalErr.Error()))
		if errErr != nil {
			log.Println(errErr.Error())
		}
	}
	// TODO respond to requester

	h.Unlock()
	w.WriteHeader(http.StatusOK)
	_ , writeErr := w.Write(jsonBytes)
	if writeErr != nil {
		log.Println(writeErr.Error())
		w.WriteHeader(http.StatusInternalServerError)
		_ , errErr := w.Write([]byte(writeErr.Error()))
		if errErr != nil {
			log.Println(errErr.Error())
		}
	}
}

func (h *linkedoutHandlers) post(w http.ResponseWriter, r *http.Request) {
	// TODO  receive request

	postMessage := h.postMongoDB(strings.TrimLeft(r.URL.Path,"/"))
	h.Lock()
	jsonBytes, marshalErr := json.Marshal(postMessage)
	if marshalErr != nil {
		log.Println(marshalErr.Error())
		w.WriteHeader(http.StatusInternalServerError)
		_ , errErr := w.Write([]byte(marshalErr.Error()))
		if errErr != nil {
			log.Println(errErr.Error())
		}
	}
	// TODO respond to requester

	h.Unlock()
	w.WriteHeader(http.StatusOK)
	_ , writeErr := w.Write(jsonBytes)
	if writeErr != nil {
		log.Println(writeErr.Error())
		w.WriteHeader(http.StatusInternalServerError)
		_ , errErr := w.Write([]byte(writeErr.Error()))
		if errErr != nil {
			log.Println(errErr.Error())
		}
	}
}