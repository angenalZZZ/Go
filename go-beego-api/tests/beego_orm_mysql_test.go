package test

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

// 1. 生成数据库实体models - DbFirst 工具 - 从数据库生成 models、routers、controllers
// > bee generate go-beego-api [-tables=""] [-driver=mysql] [-conn="root:123456@tcp(127.0.0.1:3306)/AppAuth?charset=utf8"] [-level=3]

// 2. 数据库引擎 beego/orm
var dbo orm.Ormer

const (
	mysqlDriverName = "mysql"
	mysqlConn       = "mysqlconn"
)

// 3. 数据库连接客户端初始化
func initBeegoOrmMysqlTest() {
	// 检查配置文件
	c := beego.AppConfig.String(mysqlConn)

	// 默认使用别名为default的数据库
	const aliasName = "default"
	_ = orm.RegisterDriver(mysqlDriverName, orm.DRMySQL)
	if err := orm.RegisterDataBase(aliasName, mysqlDriverName, c); err != nil {
		orm.DebugLog.Fatal("CONF " + mysqlConn + " NOT FOUND")
	}
	dbo = orm.NewOrm()
	_ = dbo.Using(aliasName)
}

// 测试: Beego Orm
func TestBeegoOrm(t *testing.T) {
	initBeegoOrmMysqlTest()
	t.Logf("DBStats: %v\n", dbo.DBStats())
}
