package conf

import (
	_ "github.com/angenalZZZ/Go/go-beego-api/routers"
	"github.com/astaxie/beego"
	"path/filepath"
	"runtime"
)

// 测试网址
var TestUrl = "http://localhost"

// 测试初始化
func TestInit() {
	_, file, _, _ := runtime.Caller(0)
	path, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	//fmt.Println(path)
	beego.TestBeegoInit(path)
	// 测试端口 RunMode = "test"
	TestUrl += ":" + beego.AppConfig.String("httpport")
	//fmt.Println(TestUrl)
}
