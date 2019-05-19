package test

import (
	"github.com/angenalZZZ/Go/go-beego-api/conf"
	_ "github.com/angenalZZZ/Go/go-beego-api/routers"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	. "github.com/smartystreets/goconvey/convey"
)

func init() {
	conf.TestInit()
}

// 方式一 (需要先手动运行 bee run)
// 端点测试：controllers/default
func TestDefault(t *testing.T) {
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

// 方式二 (全自动 Serve HTTP)
// 端点测试：controllers/user
func TestGet(t *testing.T) {
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
