package conf

import (
	_ "github.com/angenalZZZ/Go/go-beego-api/routers"
	"github.com/astaxie/beego"
)

// 测试网址
var TestUrl = "http://localhost"

// 测试初始化
func TestInit() {
	beego.TestBeegoInit(GetAppPath())

	// 测试端口 RunMode = "test"
	TestUrl += ":" + beego.AppConfig.String("httpport")
}
