package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func main(){
	router := mux.NewRouter()

	router.HandleFunc("/home", home)

	http.ListenAndServe(":8080", router)

}

func home(resp http.ResponseWriter, req *http.Request)  {
	fmt.Println("Request landed in home..!")
	fmt.Fprintf(resp, "Welcome to the HomePage!")
}
