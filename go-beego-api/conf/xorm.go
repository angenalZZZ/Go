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
	// 打印SQL执行语句 (Debug模式 帮助查看ORM与SQL执行的对照关系)
	db.ShowSQL(true)
	// 打印SQL执行时长 (Debug模式 帮助SQL执行的性能优化)
	db.ShowExecTime(true)

	// 性能优化的时候才考虑，加上本机的SQL缓存
	cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
	db.SetDefaultCacher(cacher)

	// 中国时区
	db.SetTZLocation(ChinaTimeLocation)

	// 设置空闲连接池中的最大连接数
	db.SetMaxIdleConns(10)
	// 设置数据库的最大打开连接数
	db.SetMaxOpenConns(100)
	// 设置连接可重用的最长时间
	db.SetConnMaxLifetime(600)
}
