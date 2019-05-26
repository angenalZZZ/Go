package postman

import (
	"github.com/appleboy/gofight"
	"github.com/astaxie/beego"
	"github.com/go-openapi/strfmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

const userAgent string = "PostmanRuntime"

// 测试: 当前UTC时间
func Test_CurrentUTCTime(t *testing.T) {
	url, token := "/time/now", "20429f83-8452-4568-9725-9f49f836eb02"

	a := assert.New(t)
	r := gofight.New()

	r.GET(url).SetHeader(gofight.H{
		"User-Agent":    userAgent,
		"Postman-Token": token,
	}).Run(InitHttpBasic(), func(r gofight.HTTPResponse, q gofight.HTTPRequest) {
		a.Equal(ua, q.Header.Get("User-Agent"))
		a.Equal(http.StatusOK, r.Code)
		if _, e := strfmt.ParseDateTime(r.Body.String()); e != nil {
			a.Error(e)
		}
	})
}

// 测试: 当前时间戳
func Test_CurrentTimestamp(t *testing.T) {
	url, token := "/time/now", "20429f83-8452-4568-9725-9f49f836eb02"

	b := beego.NewControllerRegister()
	b.Add(url, &MyController{}, "get:BeegoApiCurrentTimestamp")

	a := assert.New(t)
	r := gofight.New()

	r.GET(url).SetDebug(true).SetHeader(gofight.H{
		"User-Agent":    userAgent,
		"Postman-Token": token,
	}).Run(b, func(r gofight.HTTPResponse, q gofight.HTTPRequest) {
		a.Equal(http.StatusOK, r.Code)
		if _, e := strfmt.ParseDateTime(r.Body.String()); e != nil {
			a.Error(e)
		}
	})
}
