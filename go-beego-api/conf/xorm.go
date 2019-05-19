package conf

import (
	"github.com/astaxie/beego"
	"github.com/xormplus/xorm"
)

// 数据库连接客户端初始化 mssql db
func InitDbForXorm(driverName, config string) (db *xorm.Engine) {
	// 检查配置 from beego.AppConfig
	c := beego.AppConfig.String(config)
	if c == "" {
		panic("beego.AppConfig \t\t" + config + " NOT FOUND")
	}

	var e error
	// 原版方式创建引擎
	db, e = xorm.NewEngine(driverName, c)
	// 也可以针对特定数据库快捷创建
	//db, err = xorm.NewPostgreSQL(conn)
	//db, err = xorm.NewSqlite3(conn)

	// 数据库连接异常
	if e != nil {
		db.Logger().Errorf("CONF DATABASE\t\t%s\n\t\t%v", c, e)
	} else if e = db.Ping(); e != nil {
		db.Logger().Errorf("PING DATABASE\t\t%s\n\t\t%v", c, e)
	} else {
		db.Logger().Infof("PING DATABASE PASS\t\t%s", c)
		// 数据库实例配置信息
		ConfigDbForXorm(db)
	}
	return
}

// 数据库实例配置信息
func ConfigDbForXorm(db *xorm.Engine) {
	// 输出SQL执行语句
	db.ShowSQL(true)
	// 输出SQL执行时长
	db.ShowExecTime(true)

	// 设置空闲连接池中的最大连接数
	db.SetMaxIdleConns(10)
	// 设置数据库的最大打开连接数
	db.SetMaxOpenConns(100)
	// 设置连接可重用的最长时间
	db.SetConnMaxLifetime(600)
}
