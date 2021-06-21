package main

import (
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

const (
	AWS_S3_REGION = "us-east-1"
	AWS_S3_BUCKET = "cf-templates-1u4474mm6gzt-us-east-1"
)

var sess = connectAWS()

func connectAWS() *session.Session {
	sess, err := session.NewSession(
		&aws.Config{Region: aws.String(AWS_S3_REGION)})
	if err != nil {
		panic(err)
	}
	return sess
}

func main() {

	http.HandleFunc("/upload/", handlerUpload) // Upload
	http.HandleFunc("/get/", handlerDownload)  // Get the file
	log.Fatal(http.ListenAndServe(":8080", nil))
}
