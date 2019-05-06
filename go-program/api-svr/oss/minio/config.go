package minio

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/minio/minio-go"
)

// https://docs.min.io/docs/golang-client-api-reference
var (
	Cli *minio.Client

	ssl             = false // 使用ssl
	endpoint        = "127.0.0.1:9000"
	accessKeyID     = "MOS5R9S20018WBQCMQ9W"
	secretAccessKey = "FBsAnunToGHqWzuNb4ku+I8TNegVYXf0iihYTYpJ"
)

func init() {
	// 初使化 MinIO Client 对象
	var err error
	if Cli, err = minio.New(endpoint, accessKeyID, secretAccessKey, ssl); err != nil {
		log.Fatalln(err)
		return
	}

	// Add custom application details to User-Agent
	Cli.SetAppInfo("api", "1.0.0")

	// Enables HTTP tracing
	//Cli.TraceOn(os.Stderr)

	http.HandleFunc("/token/minio", WebLogin)
	http.HandleFunc("/minio/upload", Upload)
}
