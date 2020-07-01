package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type linkedoutHandlers struct {
	welcomeHandlers WelcomeJSON
}

func (h *linkedoutHandlers) get(w http.ResponseWriter, r *http.Request) {
	
	getMessage := WelcomeJSON{}
	// TODO  receive request

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
