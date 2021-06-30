package main

import (
	"encoding/json"
	//"log"
	//"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)
var v interface{}

func getCrypto(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	for{
		v = connectAPI("wss://api.hitbtc.com/api/2/ws/", "getCurrencies", "")
		time.Sleep(time.Second*5)
	}
	json.NewEncoder(w).Encode(v)
}

func getSymbols(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	for {
		v = connectAPI("wss://api.hitbtc.com/api/2/ws/", "getSymbols", "")
		time.Sleep(time.Second*5)
	}
	json.NewEncoder(w).Encode(v)
}

func getTrades(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "applicaiton/json")
	for {
		v = connectAPI("wss://api.hitbtc.com/api/2/ws/", "getTrades", "ETHBTC")
		time.Sleep(time.Second*5)
	}
	json.NewEncoder(w).Encode(v)
}
func main() {


	r := mux.NewRouter()

	r.HandleFunc("/getCrypto", getCrypto).Methods("GET")
	r.HandleFunc("/getSymbols", getSymbols).Methods("GET")
	r.HandleFunc("/getTrades", getTrades).Methods("GET")
	http.ListenAndServe(":8080", r)
}