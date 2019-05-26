package test

import (
	"github.com/angenalZZZ/Go/go-beego-api/conf"
	"github.com/angenalZZZ/Go/go-beego-api/controllers"
	_ "github.com/angenalZZZ/Go/go-beego-api/routers"
	"github.com/appleboy/gofight"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	. "github.com/smartystreets/goconvey/convey"
)

// 初始化
func init() {
	conf.TestInit()
}

// 方式一 (测试前，需要先运行 bee run)
// 端点测试：controllers/default
func TestFunction1(t *testing.T) {
	req := httplib.Get(conf.TestUrl + "/")
	res, err := req.Response()
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()
	ctn, _ := ioutil.ReadAll(res.Body)
	t.Logf("output result\n\n%s\n\n", ctn)
	if string(ctn) != beego.BConfig.AppName {
		t.Fatal("test result is not: " + beego.BConfig.AppName)
	}
}

// 方式二 (默认)
// 端点测试：controllers/user
func TestFunction2(t *testing.T) {
	r, _ := http.NewRequest("GET", "/v1/user/", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	//beego.Trace("testing", "TestGet", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			_, _ = Printf("\n1 output result\n\n%s\n\n", w.Body)
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}

// 方式三 (快捷)
// 端点测试：controllers/default
func TestFunction(t *testing.T) {
	const url = "/"
	b := beego.NewControllerRegister()
	b.Add(url, &controllers.DefaultController{}, "get:Get")

	a := assert.New(t)
	r := gofight.New()

	r.GET(url).SetDebug(true).Run(b, func(r gofight.HTTPResponse, q gofight.HTTPRequest) {
		a.Equal(http.StatusOK, r.Code)
		a.Equal(beego.BConfig.AppName, r.Body.String())
	})
}
