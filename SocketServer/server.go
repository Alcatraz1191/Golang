package main

import (
	"encoding/json"
	//"log"
	//"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func getCrypto(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	v := connectAPI("wss://api.hitbtc.com/api/2/ws/", "getCurrencies", "")
	json.NewEncoder(w).Encode(v)
}

func getSymbols(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	v := connectAPI("wss://api.hitbtc.com/api/2/ws/", "getSymbols", "")
	json.NewEncoder(w).Encode(v)
}

func getTrades(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "applicaiton/json")
	v := connectAPI("wss://api.hitbtc.com/api/2/ws/", "getTrades", "ETHBTC")
	json.NewEncoder(w).Encode(v)
}


func main() {

	r := mux.NewRouter()

	r.HandleFunc("/getCrypto", getCrypto).Methods("GET")
	r.HandleFunc("/getSymbols", getSymbols).Methods("GET")
	r.HandleFunc("/getTrades", getTrades).Methods("GET")
	http.ListenAndServe(":8080", r)
}