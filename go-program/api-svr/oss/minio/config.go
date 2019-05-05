package minio

import (
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/minio/minio-go"
)

// https://docs.min.io/docs/golang-client-api-reference
var Cli *minio.Client

const (
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
}

// 获取认证Token
func WebLogin() string {
	resp, err := http.Post("http://"+endpoint+"/minio/webrpc", "application/json", struct{ io.Reader }{strings.NewReader("")})
	if err != nil {
		return err.Error()
	}
	defer resp.Body.Close()
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err.Error()
	}
	return string(buf) // json
}
