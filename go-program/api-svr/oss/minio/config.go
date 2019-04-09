package minio

import (
	"log"

	"github.com/minio/minio-go"
)

// https://docs.min.io/docs/golang-client-api-reference
var minioClient *minio.Client

type PutObject struct {
	Location               string
	BucketName, ObjectName string
	FilePath, FileType     string
}

func init() {
	endpoint := "127.0.0.1:9000"
	accessKeyID := "MOS5R9S20018WBQCMQ9W"
	secretAccessKey := "FBsAnunToGHqWzuNb4ku+I8TNegVYXf0iihYTYpJ"

	// Initialize minio client object.
	if client, err := minio.New(endpoint, accessKeyID, secretAccessKey, false); err != nil {
		log.Fatalln(err)
	} else {
		minioClient = client
	}
}
