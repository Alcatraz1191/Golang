package main

import (
	"net/http"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func handlerDownload(w http.ResponseWriter, r *http.Request) {
	filename := strings.Replace(r.URL.Path, "/get/", "", 1)
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
		return
	}

	downloader := s3manager.NewDownloader(sess)
	_, err = downloader.Download(f, &s3.GetObjectInput{
		Bucket: aws.String(AWS_S3_BUCKET),
		Key:    aws.String(filename),
	})
	if err != nil {
		panic(err)
		return
	}
}
