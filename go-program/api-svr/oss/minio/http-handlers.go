package minio

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"mime"
	"net/http"
	"time"

	"github.com/minio/minio-go"

	"github.com/angenalZZZ/Go/go-program/api-svr/cors"
	"github.com/angenalZZZ/Go/go-program/api-svr/jsonp"
)

// File Upload Model
type PutObject struct {
	Location               string
	BucketName, ObjectName string
	FilePath, FileType     string
}

// File Upload Http Handle
func Upload(w http.ResponseWriter, r *http.Request) {
	if cors.Cors(w, r, []string{http.MethodPost}) {
		return
	}

	var p *PutObject
	defer r.Body.Close()
	if e := json.NewDecoder(r.Body).Decode(p); e != nil {
		jsonp.Error(e).Error(w, r)
		return
	}

	// 当BucketName不存在时，自动创建Bucket
	if exists, e := Cli.BucketExists(p.BucketName); e != nil {
		jsonp.Error(e).Error(w, r)
		return
	} else if exists == false {
		if e = Cli.MakeBucket(p.BucketName, p.Location); e != nil {
			jsonp.Error(e).Error(w, r)
			return
		}
	}

	// 获取上传的临时文件-表单提交file&type
	if filePath, contentType, e := getUploadedTempFile(w, r, 10*1024*1024); e != nil {
		jsonp.Error(e).Error(w, r)
		return
	} else {
		p.FilePath, p.FileType = filePath, contentType
	}

	// 保存网盘-超时设置10s
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if size, e := Cli.FPutObjectWithContext(ctx, p.BucketName, p.ObjectName, p.FilePath, minio.PutObjectOptions{ContentType: p.FileType}); e != nil {
		jsonp.Error(e).Error(w, r)
	} else {
		jsonp.Success(jsonp.Data{"data": p.ObjectName, "size": size}).OK(w, r)
	}
}

// 获取上传的临时文件-表单提交file&type
func getUploadedTempFile(w http.ResponseWriter, r *http.Request, maxUploadSize int64) (filePath string, contentType string, err error) {
	filePath, contentType, err = "", "", nil

	// validate file size
	r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)
	if e := r.ParseMultipartForm(maxUploadSize); e != nil {
		err = errors.New(fmt.Sprintf("文件大小限制：%dM", maxUploadSize/(1024*1024)))
		return
	}

	// parse and validate file and post parameters
	file, _, e := r.FormFile("file")
	if e != nil {
		err = errors.New("未发现文件上传")
		return
	}
	defer file.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		err = errors.New("无效的文件上传")
		return
	}

	// check file type, detect content type only needs the first 512 bytes
	contentType = http.DetectContentType(fileBytes)
	switch contentType {
	case "image/jpeg", "image/jpg":
	case "image/gif", "image/png":
	case "application/pdf":
		break
	default:
		err = errors.New("无效的文件类型:" + contentType)
		return
	}

	// check post file type, .jpg .gif .png
	fileType := r.PostFormValue("type")
	fileEndings, e := mime.ExtensionsByType(fileType)
	if e != nil {
		err = errors.New("未发现文件上传类型")
		return
	}

	//b := make([]byte, 6)
	//rand.Read(b)
	//fileName := fmt.Sprintf("%d%x", time.Now().Unix(), b)
	//newPath := filepath.Join(os.TempDir(), fileName+fileEndings[0])
	//fmt.Printf("FileType: %s, File: %s\n", fileType, newPath)

	// write file
	//newFile, e := os.Create(newPath)
	//if e != nil {
	//	err = errors.New("无法创建临时文件")
	//	return
	//}

	newFile, e := ioutil.TempFile("", "*"+fileEndings[0])
	filePath = newFile.Name()
	defer newFile.Close() // idempotent, okay to call twice
	if _, err := newFile.Write(fileBytes); err != nil || newFile.Close() != nil {
		err = errors.New("无法创建临时文件")
	}
	return
}
