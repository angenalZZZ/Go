package minio

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/minio/minio-go"

	"github.com/angenalZZZ/Go/go-program/api-svr/cors"
	"github.com/angenalZZZ/Go/go-program/api-svr/jsonp"
)

// File Upload
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
	if exists, e := minioClient.BucketExists(p.BucketName); e != nil {
		jsonp.Error(e).Error(w, r)
		return
	} else if exists == false {
		if e = minioClient.MakeBucket(p.BucketName, p.Location); e != nil {
			jsonp.Error(e).Error(w, r)
			return
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if size, e := minioClient.FPutObjectWithContext(ctx, p.BucketName, p.ObjectName, p.FilePath, minio.PutObjectOptions{ContentType: p.FileType}); e != nil {
		jsonp.Error(e).Error(w, r)
		return
	} else {
		jsonp.Success(jsonp.Data{"data": p.ObjectName, "size": size}).OK(w, r)
	}

}
