package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

func newLinkedoutHandlers() *linkedoutHandlers {
	return &linkedoutHandlers{
		data: map[string]WelcomeJSON{},
	}
}

type linkedoutHandlers struct {
	sync.Mutex
	welcomeHandlers WelcomeJSON
	contactHandlers ContactJSON
}

func (h *linkedoutHandlers) get(w http.ResponseWriter, r *http.Request) {
	// TODO  receive request

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
