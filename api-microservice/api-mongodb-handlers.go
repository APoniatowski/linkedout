package main

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
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

	switch col {
	case "welcome":
		// as only 1 thing will be returned, no need to create a slice/array
		cursor, err := collection.Find(ctx, bson.D{})
		if err != nil {
			log.Println(err)
		}
		defer cursor.Close(ctx)
		return cursor
	case "contact":
		// as only 1 thing will be returned, no need to create a slice/array
		cursor, err := collection.Find(ctx, bson.D{})
		if err != nil {
			log.Println(err)
		}
		defer cursor.Close(ctx)
		return cursor
	default:
		log.Println("404: Page not found")
		return sendError
	}
}

func (h *linkedoutHandlers) postMongoDB(col string) string {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	// Connect to MongoDB
	timeout := 5 * time.Second
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
		if welcomeCheck > 0 {
			// convert struct/json to bson
			filterMessage := collection.FindOne(ctx, "message")
			updateMessage := bson.M{"$set": bson.M{"message": h.welcomeHandlers.Message}}
			// Set FindOneAndUpdateOptions
			upsert := true
			after := options.After
			opt := options.FindOneAndUpdateOptions{
				ReturnDocument: &after,
				Upsert:         &upsert,
			}
			// Find and update it
			res := collection.FindOneAndUpdate(ctx, filterMessage, updateMessage, &opt)
			if res.Err() != nil {
				log.Fatal(res.Err())
				return sendError
			}
			return sendSuccess
		} else {
			res, err := collection.InsertOne(ctx, h.welcomeHandlers)
			if err != nil {
				log.Fatal(err)
				return sendError
			}
			id := res.InsertedID
			return sendSuccess + id.(string)
		}
	case "contact":
		welcomeCheck, checkErr := collection.EstimatedDocumentCount(ctx)
		if checkErr != nil {
			log.Fatal(err)
		}
		if welcomeCheck > 0 {
			// convert struct/json to bson
			filterMessage := collection.FindOne(ctx, "email")
			updateMessage := bson.M{"$set": bson.M{"email": h.contactHandlers.Email,
				"linkedinurl": h.contactHandlers.LinkedinURL,
				"phone": h.contactHandlers}}
			// Set FindOneAndUpdateOptions
			upsert := true
			after := options.After
			opt := options.FindOneAndUpdateOptions{
				ReturnDocument: &after,
				Upsert:         &upsert,
			}
			// Find and update it
			res := collection.FindOneAndUpdate(ctx, filterMessage, updateMessage, &opt)
			if res.Err() != nil {
				log.Fatal(res.Err())
				return sendError
			}
			return sendSuccess
		} else {
			res, err := collection.InsertOne(ctx, h.contactHandlers)
			if err != nil {
				log.Fatal(err)
				return sendError
			}
			id := res.InsertedID
			return sendSuccess + id.(string)
		}
	default:
		return sendError
	}

}