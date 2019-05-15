package main

import (
	_ "github.com/angenalZZZ/Go/go-beego-api/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	//_ "github.com/lib/pq"
	//_ "github.com/mattn/go-sqlite3"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	_ = orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("mysqlconn"))
	//o := orm.NewOrm() // 默认使用别名为default的数据库
	//_ = o.Using("default") // Ormer interface
	DbFirst_Bee_Generate() // Db First
}

// Db First 从数据库生成 models、routers、controllers
// > bee generate go-beego-api [-tables=""] [-driver=mysql] [-conn="root:123456@tcp(127.0.0.1:3306)/AppAuth?charset=utf8"] [-level=3]
func DbFirst_Bee_Generate() {
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
