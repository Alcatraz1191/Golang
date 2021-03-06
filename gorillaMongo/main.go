package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"gorillaMongo/helper"
	"gorillaMongo/models"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Connection mongoDB with helper class
var collection = helper.ConnectDB()
func main(){
	//Init Router
	r := mux.NewRouter()

	//arrange route
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	//set port
	log.Fatal(http.ListenAndServe(":8080", r))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	//set header
	w.Header().Set("Content-Type", "application/json")

	//create a book array
	var books []models.Book

	//bson.M{}, pass empty filter to get all data
	cur, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		helper.GetError(err, w)
	}
	// defer closing cur
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		// create variable to decode document into
		var book models.Book
		err := cur.Decode(&book)
		if err != nil {
			log.Fatal(err)
		}

		//add item to our array
		books = append(books, book)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(books)

}

func getBook(w http.ResponseWriter, r *http.Request) {
	//set header
	w.Header().Set("Content-Type", "application/json")

	var book models.Book

	//get params with mux
	var params = mux.Vars(r)

	//string to ObjectID
	id, _ := primitive.ObjectIDFromHex(params["id"])

	//create filter
	filter := bson.M{"_id":id}

	err := collection.FindOne(context.TODO(), filter).Decode(&book)

	if err != nil {
		helper.GetError(err, w)
		return
	}
	json.NewEncoder(w).Encode(book)

}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book models.Book

	//decode body
	_= json.NewDecoder(r.Body).Decode(&book)

	//insert our book model
	result, err := collection.InsertOne(context.TODO(), book)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var params = mux.Vars(r)

	//get id from parameters
	id, _ := primitive.ObjectIDFromHex(params["id"])

	var book models.Book

	//create filter
	filter := bson.M{"_id": id}

	//Read update model from body
	_ = json.NewDecoder(r.Body).Decode(&book)

	update := bson.D{
		{"$set", bson.D{
			{"isbn", book.Isbn},
			{"title", book.Title},
			{"author", bson.D{
				{"firstname" , book.Author.FirstName},
				{"lastname", book.Author.LastName},
			}},
		}},
	}
	err := collection.FindOneAndUpdate(context.TODO(), filter , update).Decode(&book)

	if err != nil{
		helper.GetError(err, w)
		return
	}

	book.ID = id

	json.NewEncoder(w).Encode(book)
}

func deleteBook (w http.ResponseWriter, r *http.Request) {
	//set header
	w.Header().Set("Content-Type", "application/json")

	//get params
	params := mux.Vars(r)

	//string to primitive object.ID
	id, _ := primitive.ObjectIDFromHex(params["id"])

	//prepare filter
	filter := bson.M{"_id": id}

	deleteResult, err := collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(deleteResult)
}