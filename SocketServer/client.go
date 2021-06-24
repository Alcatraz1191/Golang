package main

import (
	"fmt"
	"time"

	"github.com/gorilla/websocket"
)


type Response struct {
	Jsonrpc string      `json:"jsonrpc"`
	ID      interface{} `json:"id"`
	Result  []struct {
		ID                 string `json:"id"`
		FullName           string `json:"fullName"`
		Crypto             bool   `json:"crypto"`
		PayinEnabled       bool   `json:"payinEnabled"`
		PayinPaymentID     bool   `json:"payinPaymentId"`
		PayinConfirmations int    `json:"payinConfirmations"`
		PayoutEnabled      bool   `json:"payoutEnabled"`
		PayoutIsPaymentID  bool   `json:"payoutIsPaymentId"`
		TransferEnabled    bool   `json:"transferEnabled"`
		Delisted           bool   `json:"delisted"`
		PayoutFee          string `json:"payoutFee,omitempty"`
	} `json:"result"`
}

type Request struct {
	Method string `json:"method"`
	Params struct{
		Symbol string `json:"symbol"`
	} `json:"params"`
}

type TradeResponse struct {
	Jsonrpc string      `json:"jsonrpc"`
	ID      interface{} `json:"id"`
	Result  Result      `json:"result"`
}
type Data struct {
	ID        int       `json:"id"`
	Price     string    `json:"price"`
	Quantity  string    `json:"quantity"`
	Side      string    `json:"side"`
	Timestamp time.Time `json:"timestamp"`
}
type Result struct {
	Data []Data `json:"data"`
}

func connectAPI(connString string, methodString string, paramSymbol string) interface{} {
	var (
		v Response
		m Request
		dialer websocket.Dialer
		t TradeResponse
	)
	m.Method = methodString
	m.Params.Symbol = paramSymbol
	fmt.Println(m)
	conn, _, err := dialer.Dial(connString, nil)
	if err != nil {
		panic(err)
	}
	err = conn.WriteJSON(m)
	if err != nil {
		panic(err)
	}
	fmt.Println(m)
	if paramSymbol == ""{
		err = conn.ReadJSON(&v)
	} else {
		err = conn.ReadJSON(&t)
		return t
	}

	if err != nil {
		panic(err)
	}
	return v
}
