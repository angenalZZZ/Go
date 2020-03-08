package sign_md5

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

var AppSecret string

// MD5 组合加密
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
	value, ok := config.ApiAuthConfig[ak]
	if ok {
		AppSecret = value["md5"]
	} else {
		return nil, errors.New("ak Error")
	}

	if debug == "1" {
		currentUnix := time.Now().Unix()
		req.Set("ts", strconv.FormatInt(currentUnix, 10))
		res := map[string]string{
			"ts": strconv.FormatInt(currentUnix, 10),
			"sn": createSign(req),
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
	if sn == "" || sn != createSign(req) {
		return nil, errors.New("sn Error")
	}

	return nil, nil
}

// 创建签名
func createSign(params url.Values) string {
	// 自定义 MD5 组合
	return f.CryptoMD5(AppSecret + createEncryptStr(params) + AppSecret)
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
