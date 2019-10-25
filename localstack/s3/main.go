package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"log"
	"os"
)

const (
	REGION       = "us-east-1"
	BUCKET       = "demo-bucket"
	S3Dictionary = "report-exports"
	Filename = "test.pdf"
	S3Endpoint = "http://localhost:4572"
)

func main() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(REGION),
		Endpoint: aws.String(S3Endpoint),
		S3ForcePathStyle: aws.Bool(true),
	})
	if err != nil {
		log.Printf("Error when establish aws session: %s\n", err)
		return;
	}

    s3uploader := s3manager.NewUploader(sess)

    fileReader, err := os.Open(Filename)
    if err != nil {
    	log.Printf("Error opening file to upload to S3: %s", Filename)
		return;
	}
    defer fileReader.Close()


    _, err = s3uploader.Upload(&s3manager.UploadInput{
    	Bucket: aws.String(BUCKET),
    	Key: aws.String(fmt.Sprintf("%s/%s", S3Dictionary, Filename)),
    	Body: fileReader,
	})
    if err != nil {
		log.Printf("Error when uploading document: %s\n", err)
		return;
	}
    log.Printf("Upload file to s3 succesfully")
}
