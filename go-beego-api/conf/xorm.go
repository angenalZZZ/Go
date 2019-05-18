package conf

import (
	"github.com/xormplus/xorm"
)

// 数据库连接客户端初始化
func InitDbForXorm(path string) (db *xorm.Engine) {
	// 检查配置文件
	con, err := GetAppConfig("mssqlconn", path)
	if err != nil {
		db.Logger().Error(err)
	}

	// 原版方式创建引擎
	db, err = xorm.NewEngine("mssql", con.(string))
	// 也可以针对特定数据库快捷创建
	//db, err = xorm.NewPostgreSQL(conn)
	//db, err = xorm.NewSqlite3(conn)

	// 数据库连接异常
	if err != nil {
		db.Logger().Errorf("CONF DATABASE\t\t%s\n\t\t%v", con, err)
	} else if err = db.Ping(); err != nil {
		db.Logger().Errorf("PING DATABASE\t\t%s\n\t\t%v", con, err)
	} else {
		db.Logger().Infof("PING DATABASE PASS\t\t%s", con)
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
