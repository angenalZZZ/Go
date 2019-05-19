package main

import (
	_ "github.com/angenalZZZ/Go/go-beego-api/routers"
	"github.com/astaxie/beego"
)

func main() {
	// 开发模式
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		// 运行App带API说明文档：bee run -gendoc=true -downdoc=true
		// 放在此目录,swagger下载地址 github.com/beego/swagger/releases
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
		// API 增加 CORS 支持
		// ctx.Output.Header("Access-Control-Allow-Origin", "*")
	}

	// 运行App
	beego.Run()

	// 生成代码工具
	// 1.从数据库一键生成 model、controller、view
	// bee generate scaffold [scaffold-name] [-fields=""] [-driver=mysql] [-conn="root:@tcp(127.0.0.1:3306)/test"]
	// 2.生成一个 model
	// bee generate model [model-name] [-fields=""]
	// 3.生成一个 controller
	// bee generate controller [controller-file]
	// 4.生成一个 view CRUD in path
	// bee generate view [view-path]
	// 5.生成数据库迁移与结构变更 file for: https://beego.me/docs/install/bee.md#migrate-命令
	// bee generate migration [migration-file] [-fields=""]
	// 6.生成API说明文档 swagger.json & swagger.yml
	// bee generate docs
	// 7.生成数据库models - DbFirst
	// bee generate [app-code] [-tables=""] [-driver=mysql] [-conn="root:@tcp(127.0.0.1:3306)/test"] [-level=3]
	//   -level:  [1 | 2 | 3], 1 = models; 2 = models,controllers; 3 = models,controllers,router

	// 生成Dockerfile文件来实现docker应用
	// bee dockerize -image="library/golang:1.11.5" -expose=80
}
