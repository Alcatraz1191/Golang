package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

type JSONData struct {
	ID      interface{} `json:"id"`
	Jsonrpc string      `json:"jsonrpc"`
	Result  struct {
		Data []struct {
			ID        int64  `json:"id"`
			Price     string `json:"price"`
			Quantity  string `json:"quantity"`
			Side      string `json:"side"`
			Timestamp string `json:"timestamp"`
		} `json:"data"`
	} `json:"result"`
}

type TokenStore struct {
	TokenType    string `json:"token_type"`
	ExpiresIn    string `json:"expires_in"`
	ExtExpiresIn string `json:"ext_expires_in"`
	ExpiresOn    string `json:"expires_on"`
	NotBefore    string `json:"not_before"`
	Resource     string `json:"resource"`
	AccessToken  string `json:"access_token"`
}

type AccessCode struct {
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
	ExpiresIn    string `json:"expires_in"`
	ExpiresOn    string `json:"expires_on"`
	NotBefore    string `json:"not_before"`
	Resource     string `json:"resource"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	IDToken      string `json:"id_token"`
}

var (
	TENANT_ID = "b60a4dea-e2d5-4d1f-b023-fd03ef62c77b"    
 	APPLICATION_ID = "fbe28bd4-3b85-47c1-805e-856ec7a7b92a" 
	REDIRECT_URL_ENCODE = "http%3A%2F%2Flocalhost%3A8080%0A"
	REDIRECT_URL_DECODE = "http://localhost:8080"
	CLIENT_SECRET = "lyQUC~8-yI~dV48.wvnK003Rmj-ApxDnC2"
	RESOURCE_URL = "https://datalake.azure.net/"
)
var tokenresp *TokenStore


func getToken(){
	val := url.Values{}
	val.Add("grant_type", "client_credentials")
	val.Add("client_id", APPLICATION_ID)
	val.Add("client_secret", CLIENT_SECRET)
	val.Add("resource", RESOURCE_URL)
	//w.Header().Set("Content-Type", "x-www-form-urlencoded")
	urlstr := "https://login.microsoftonline.com/b60a4dea-e2d5-4d1f-b023-fd03ef62c77b/oauth2/token"
	client := &http.Client{}

	resp, _ := client.PostForm(urlstr, val)
	fmt.Println(resp.Status)
	fmt.Println(resp.Body)
   json.NewDecoder(resp.Body).Decode(&tokenresp)
   fmt.Println(tokenresp.AccessToken)
}

func main(){
	r := mux.NewRouter()
	getToken()
	r.HandleFunc("/openFile/{filename}", openFile).Methods("GET")
	r.HandleFunc("/uploadFile", uploadFile).Methods("GET")
	http.ListenAndServe(":8080", r)
}

func openFile(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	w.Header().Set("Authorization", fmt.Sprintf("Bearer %v", tokenresp.AccessToken))
	w.Header().Set("Content-Type", "application/json")
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://mystorage1.azuredatalakestore.net/webhdfs/v1/test/"+params["filename"]+"/?op=OPEN", nil)
	req.Header.Set("Authorization", "Bearer "+tokenresp.AccessToken)
	resp, _ := client.Do(req)
	respData, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(resp.Body)
	claims := jwt.MapClaims{}

	var p *jwt.Parser
	var jdata *JSONData
	_, parts, _ := p.ParseUnverified(string(respData), claims)
	s := strings.Join(parts, "")
	_ = json.Unmarshal([]byte(s), &jdata)
	json.NewEncoder(w).Encode(&jdata)
}

func uploadFile(w http.ResponseWriter, r *http.Request){
	file, err := os.Open("file")
	if err != nil {
		panic(err)
	}
	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	fi, err := file.Stat()
	if err != nil {
		panic(err)
	}
	file.Close()

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", fi.Name())
	if err != nil {
		panic(err)
	}
	part.Write(fileContents)

	err = writer.Close()
	if err != nil {
		panic(err)
	}

	req, _ := http.NewRequest("PUT", "https://mystorage1.azuredatalakestore.net/webhdfs/v1/test/file?op=CREATE", body)

	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

}
