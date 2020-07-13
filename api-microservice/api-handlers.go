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
	success = "success"
	error = "error"
)


func (h *linkedoutHandlers) getMongoDB(col string) interface{} {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	// Connect to MongoDB
	timeout := 5*time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
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
	timeout := 5*time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	collection := client.Database("linkedout").Collection(col)
	switch col {
	case "welcome":
		welcomeCheck, checkErr := collection.EstimatedDocumentCount(ctx)
		if checkErr != nil {
			log.Fatal(err)
		}
		if welcomeCheck == 0 {
			res, err := collection.InsertOne(ctx, h.welcomeHandlers)
			if err != nil {
				log.Fatal(err)
				return error
			}
			id := res.InsertedID
			return success + id.(string)
		} else {
			// convert struct/json to bson
			filterMessage := collection.FindOne(ctx,"message")
			updateMessage := bson.M{"$set": bson.M{"message":h.welcomeHandlers.Message}}
			// Set FindOneAndUpdateOptions
			upsert := true
			after := options.After
			opt := options.FindOneAndUpdateOptions{
				ReturnDocument: &after,
				Upsert:         &upsert,
			}
			// Find and update it
			res := collection.FindOneAndUpdate(ctx,filterMessage,updateMessage,&opt)
			if res.Err() != nil {
				log.Fatal(res.Err())
				return error
			}

			return success
		}
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
