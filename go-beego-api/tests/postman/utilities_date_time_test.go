package postman

import (
	"github.com/appleboy/gofight"
	"github.com/astaxie/beego"
	"github.com/go-openapi/strfmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 测试: 当前时间
func Test_CurrentUTCTime(t *testing.T) {
	a := assert.New(t)
	r := gofight.New()

	url, ua := "/time/now", "PostmanRuntime"

	r.GET(url).SetHeader(gofight.H{
		"User-Agent":    ua,
		"Postman-Token": "20429f83-8452-4568-9725-9f49f836eb02",
	}).Run(InitHttpBasic(), func(r gofight.HTTPResponse, q gofight.HTTPRequest) {
		a.Equal(ua, q.Header.Get("User-Agent"))
		a.Equal(http.StatusOK, r.Code)
		if _, e := strfmt.ParseDateTime(r.Body.String()); e != nil {
			a.Error(e)
		}
	})
}

// 测试: 当前时间
func Test_CurrentTimestamp(t *testing.T) {
	url := "/time/now"

	b := beego.NewControllerRegister()
	b.Add(url, &MyController{}, "get:BeegoApiCurrentTimestamp")

	a := assert.New(t)
	r := gofight.New()

	r.GET(url).SetDebug(true).Run(b, func(r gofight.HTTPResponse, q gofight.HTTPRequest) {
		a.Equal(http.StatusOK, r.Code)
		if _, e := strfmt.ParseDateTime(r.Body.String()); e != nil {
			a.Error(e)
		}
	})
}
