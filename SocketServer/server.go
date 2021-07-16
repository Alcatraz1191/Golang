package main

import (
	"encoding/json"
	//"log"
	//"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type GetSymbolJSON struct {
	AvgProcessingTime  string `json:"avgProcessingTime"`
	Crypto             bool   `json:"crypto"`
	Delisted           bool   `json:"delisted"`
	FullName           string `json:"fullName"`
	HighProcessingTime string `json:"highProcessingTime"`
	ID                 string `json:"id"`
	LowProcessingTime  string `json:"lowProcessingTime"`
	PayinConfirmations int64  `json:"payinConfirmations"`
	PayinEnabled       bool   `json:"payinEnabled"`
	PayinPaymentID     bool   `json:"payinPaymentId"`
	PayoutEnabled      bool   `json:"payoutEnabled"`
	PayoutFee          string `json:"payoutFee"`
	PayoutIsPaymentID  bool   `json:"payoutIsPaymentId"`
	PrecisionPayout    int64  `json:"precisionPayout"`
	PrecisionTransfer  int64  `json:"precisionTransfer"`
	TransferEnabled    bool   `json:"transferEnabled"`
}


var v interface{}

var SymbolJSON GetSymbolJSON

func getCrypto(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	for {
		v = connectAPI("wss://api.hitbtc.com/api/2/ws/", "getCurrencies", "")
		time.Sleep(time.Second * 5)
	}
	json.NewEncoder(w).Encode(v)
}

func getSymbols(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	for {
		v = connectAPI("wss://api.hitbtc.com/api/2/ws/", "getSymbols", "")
		time.Sleep(time.Second * 5)
	}
	json.NewEncoder(w).Encode(v)
}

func getTrades(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applicaiton/json")
	for {
		v = connectAPI("wss://api.hitbtc.com/api/2/ws/", "getTrades", "ETHBTC")
		time.Sleep(time.Second * 5)
	}
	json.NewEncoder(w).Encode(v)
}

func getSymbol(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applicaiton/json")

	v, _ = http.Get("http://api.hitbtc.com/api/2/public/currency")


	json.NewEncoder(w).Encode(SymbolJSON)
}
func main() {

	r := mux.NewRouter()

	r.HandleFunc("/getCrypto", getCrypto).Methods("GET")
	r.HandleFunc("/getSymbols", getSymbols).Methods("GET")
	r.HandleFunc("/getTrades", getTrades).Methods("GET")
	http.ListenAndServe(":8080", r)
}
