package models

import (
	"github.com/angenalZZZ/Go/go-beego-api/conf"

	"github.com/astaxie/beego/orm"
	"testing"

	//_ "github.com/lib/pq"
	//_ "github.com/mattn/go-sqlite3"
	_ "github.com/go-sql-driver/mysql"
)

// 1. 生成数据库实体models - DbFirst 工具 - 从数据库生成 models、routers、controllers
// > bee generate go-beego-api [-tables=""] [-driver=mysql] [-conn="root:123456@tcp(127.0.0.1:3306)/AppAuth?charset=utf8"] [-level=3]

// 2. 数据库引擎 beego/orm
var dbo orm.Ormer

// 3. 数据库连接客户端初始化
func init() {
	// 检查配置文件
	con, err := conf.GetAppConfig("mysqlconn", "../conf/app.conf")
	if err != nil {
		orm.DebugLog.Fatal(err)
	}
	// 默认使用别名为default的数据库
	const aliasName = "default"
	if err = orm.RegisterDataBase(aliasName, "mysql", con.(string)); err != nil {
		orm.DebugLog.Fatal("CONF DATABASE\t\tmysqlconn NOT FOUND")
	}
	dbo = orm.NewOrm()
	_ = dbo.Using(aliasName)
}

func TestBeegoOrm(t *testing.T) {

}
