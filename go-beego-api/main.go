package main

import (
	_ "github.com/angenalZZZ/Go/go-beego-api/routers"
	"github.com/astaxie/beego/orm"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	_ = orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("sqlconn"))
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
