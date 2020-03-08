package sign_aes

import (
	"errors"
	"fmt"
	"github.com/angenalZZZ/Go/go-gin-api/server/config"
	"github.com/angenalZZZ/Go/go-gin-api/server/util/ctx"
	"github.com/angenalZZZ/gofunc/f"
	"github.com/gin-gonic/gin"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

// AES 对称加密
func SetUp() gin.HandlerFunc {

	return func(c *gin.Context) {
		g := ctx.Wrap(c)

		sign, err := verifySign(c)

		if sign != nil {
			g.Fail("Debug Sign", sign)
			c.Abort()
			return
		}

		if err != nil {
			g.Fail(err.Error(), sign)
			c.Abort()
			return
		}

		c.Next()
	}
}

// 验证签名
func verifySign(c *gin.Context) (map[string]string, error) {
	_ = c.Request.ParseForm()
	req := c.Request.Form
	debug := strings.Join(c.Request.Form["debug"], "")
	ak := strings.Join(c.Request.Form["ak"], "")
	sn := strings.Join(c.Request.Form["sn"], "")
	ts := strings.Join(c.Request.Form["ts"], "")

	// 验证来源
	appSecret := ""
	value, ok := config.ApiAuthConfig[ak]
	if ok {
		appSecret = value["aes"]
	} else {
		return nil, errors.New("ak Error")
	}

	if debug == "1" {
		currentUnix := time.Now().Unix()
		req.Set("ts", strconv.FormatInt(currentUnix, 10))

		sn, err := createSign(req, appSecret)
		if err != nil {
			return nil, errors.New("sn Exception")
		}

		res := map[string]string{
			"ts": strconv.FormatInt(currentUnix, 10),
			"sn": string(sn),
		}
		return res, nil
	}

	// 验证过期时间
	timestamp := time.Now().Unix()
	exp, _ := strconv.ParseInt(config.AppSignExpiry, 10, 64)
	tsInt, _ := strconv.ParseInt(ts, 10, 64)
	if tsInt > timestamp || timestamp-tsInt >= exp {
		return nil, errors.New("ts Error")
	}

	// 验证签名
	if sn == "" {
		return nil, errors.New("sn Error")
	}

	decryptStr, decryptErr := f.CryptoAesCBCEncrypt([]byte(sn), []byte(appSecret), []byte(appSecret))
	if decryptErr != nil {
		return nil, errors.New(decryptErr.Error())
	}
	if string(decryptStr) != createEncryptStr(req) {
		return nil, errors.New("sn Error")
	}
	return nil, nil
}

// 创建签名
func createSign(params url.Values, appSecret string) ([]byte, error) {
	return f.CryptoAesCBCDecrypt([]byte(createEncryptStr(params)), []byte(appSecret), []byte(appSecret))
}

func createEncryptStr(params url.Values) string {
	var key []string
	var str = ""
	for k := range params {
		if k != "sn" && k != "debug" {
			key = append(key, k)
		}
	}
	sort.Strings(key)
	for i := 0; i < len(key); i++ {
		if i == 0 {
			str = fmt.Sprintf("%v=%v", key[i], params.Get(key[i]))
		} else {
			str = str + fmt.Sprintf("&%v=%v", key[i], params.Get(key[i]))
		}
	}
	return str
}
