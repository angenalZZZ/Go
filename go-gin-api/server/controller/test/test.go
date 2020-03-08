package test

import (
	"fmt"
	"github.com/angenalZZZ/Go/go-gin-api/server/util/ctx"
	"github.com/angenalZZZ/gofunc/f"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"time"
)

func Md5Test(c *gin.Context) {
	startTime := time.Now()
	appSecret := "IgkibX71IEf382PT"
	encryptStr := "param_1=xxx&param_2=xxx&ak=xxx&ts=1111111111"
	count := 1000000
	for i := 0; i < count; i++ {
		// 生成签名
		f.CryptoMD5(appSecret + encryptStr + appSecret)

		// 验证签名
		f.CryptoMD5(appSecret + encryptStr + appSecret)
	}
	ctx.Wrap(c).OK(fmt.Sprintf("%v次 - %v", count, time.Since(startTime)), nil)
}

func AesTest(c *gin.Context) {
	startTime := time.Now()
	appSecret := "IgkibX71IEf382PT"
	encryptStr := "param_1=xxx&param_2=xxx&ak=xxx&ts=1111111111"
	count := 1000000
	for i := 0; i < count; i++ {
		// 生成签名
		sn, _ := f.CryptoAesCBCDecrypt([]byte(encryptStr), []byte(appSecret), []byte(appSecret))

		// 验证签名
		_, _ = f.CryptoAesCBCDecrypt(sn, []byte(appSecret), []byte(appSecret))
	}
	ctx.Wrap(c).OK(fmt.Sprintf("%v次 - %v", count, time.Since(startTime)), nil)
}

func RsaTest(c *gin.Context) {
	startTime := time.Now()
	encryptStr := "param_1=xxx&param_2=xxx&ak=xxx&ts=1111111111"

	publicKeyPemFile, privateKeyPemFile := "rsa/public.pem", "rsa/private.pem"
	publicKeyPemBytes, _ := ioutil.ReadFile(publicKeyPemFile)
	privateKeyPemBytes, _ := ioutil.ReadFile(privateKeyPemFile)
	publicKeyEncrypt := f.NewRSAPublicKeyEncrypt(publicKeyPemBytes)
	privateKeyDecrypt := f.NewRSAPrivateKeyDecrypt(privateKeyPemBytes)

	count := 500
	for i := 0; i < count; i++ {
		// 生成签名
		sn, _ := publicKeyEncrypt.EncryptPKCS1v15([]byte(encryptStr))

		// 验证签名
		_, _ = privateKeyDecrypt.DecryptPKCS1v15(sn)
	}
	ctx.Wrap(c).OK(fmt.Sprintf("%v次 - %v", count, time.Since(startTime)), nil)
}
