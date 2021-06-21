package helper

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//connectDB: Lets you connect to mongoDB
//start with uppercase to see function when you import on other class
func ConnectDB() *mongo.Collection {

	//set client options
	ClientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017/?readPreference=primary&appname=mongodb-vscode%200.5.0&ssl=false")

	client, err := mongo.Connect(context.TODO(), ClientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection := client.Database("go_rest_api").Collection("books")

	return collection
}

//ErrorResponse error model

type ErrorResponse struct {
	StatusCode   int    `json:"status"`
	ErrorMessage string `json:"message"`
}

//GetError helper function for error model
//start with capital letter
func GetError(err error, w http.ResponseWriter) {

	log.Fatal(err.Error())
	var response = ErrorResponse{
		ErrorMessage: err.Error(),
		StatusCode:   http.StatusInternalServerError,
	}

	message, _ := json.Marshal(response)

	w.WriteHeader(response.StatusCode)
	w.Write(message)
}
