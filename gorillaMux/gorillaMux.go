package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	//"strconv"
)

type Post struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

var posts []Post

func main() {

	posts = append(posts, Post{ID: "1", Title: "My first post", Body: "The content of my first post is Nobunaga no Shinobi"})
	posts = append(posts, Post{ID: "2", Title: "My second post", Body: "This is the content of my second post"})
	r := mux.NewRouter()

	r.HandleFunc("/posts", getPosts).Methods("GET")
	r.HandleFunc("/posts/{id}", getPost).Methods("GET")
	r.HandleFunc("/posts", createPost).Methods("POST")
	r.HandleFunc("/posts/{id}", updatePost).Methods("PUT")
	r.HandleFunc("/posts/{id}", deletePost).Methods("DELETE")

	http.ListenAndServe(":8080", r)
}

func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func getPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range posts {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			break
		}
	}
	json.NewEncoder(w).Encode(&Post{})
}

func createPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var post Post

	_ = json.NewDecoder(r.Body).Decode(&post)
	posts = append(posts, post)
	json.NewEncoder(w).Encode(&post)
}

func updatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, item := range posts {
		if item.ID == params["id"] {
			posts = append(posts[:i], posts[i+1:]...)

			var post Post
			_ = json.NewDecoder(r.Body).Decode(&post)
			posts = append(posts, post)
			return
		}
	}
	json.NewEncoder(w).Encode(posts)
}

func deletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, item := range posts {
		if item.ID == params["id"] {
			posts = append(posts[:i], posts[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(posts)
}
