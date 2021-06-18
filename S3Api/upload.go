package main

import (
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func handlerUpload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)
	file, header, err := r.FormFile("file")
	if err != nil {
		panic(err)
		return
	}
	defer file.Close()

	filename := header.Filename

	uploader := s3manager.NewUploader(sess)
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(AWS_S3_BUCKET), // Bucket to be used
		Key:    aws.String(filename),      // Name of the file to be saved
		Body:   file,
	})
	if err != nil {
		panic(err)
		return
	}
}
