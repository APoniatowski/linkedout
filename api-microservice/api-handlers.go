package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	// go get go.mongodb.org/mongo-driver
)

const (
	success = "SUCCESS"
	error = "SUCCESS"
)


func (h *linkedoutHandlers) getMongoDB(col string) interface{} {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	collection := client.Database("linkedout").Collection(col)
	//fmt.Println(collection)
	return collection
}

func (h *linkedoutHandlers) postMongoDB(col string) string {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	collection := client.Database("linkedout").Collection(col)
	switch col {
	case "welcome":
		res, err := collection.InsertOne(context.TODO(), h.welcomeHandlers)
		if err != nil {
			log.Fatal(err)
			return error
		}
		id := res.InsertedID
		return success
	default:
		return error
	}



}

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
