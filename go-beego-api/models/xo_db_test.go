package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/google/uuid"
	"github.com/xormplus/xorm"
	"testing"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
)

// 1. 生成数据库实体models - use DbFirst tool - https://github.com/xo/dburl
//  > xo mysql://root:123456@127.0.0.1:3306/AppAuth?parseTime=true -o ./models
//  > xo mssql://sa:Your_password123@localhost:1434/AppAuth?parseTime=true -o ./models

// 2. 数据库引擎增强版 xorm db engine
var db *xorm.Engine

func init() {
	var err error
	if err = beego.LoadAppConfig("ini", "../conf/app.conf"); err != nil {
		panic(err)
	}

	sqlconn := beego.AppConfig.String("mssqlconn")

	// 原版方式创建引擎
	db, err = xorm.NewEngine("mssql", sqlconn)
	// 也可以针对特定数据库快捷创建
	//db, err = xorm.NewPostgreSQL(sqlconn)
	//db, err = xorm.NewSqlite3(sqlconn)

	// 数据库连接异常
	if err != nil {
		panic(sqlconn + " -> " + err.Error())
	} else {
		fmt.Printf("%s -> OK\n", sqlconn)
	}
}

func TestUUID(t *testing.T) {
	src1, src2 := uuid.New(), NewID()
	t.Logf("TestUUID32: %s \t %s", src1, src2)
}

func TestAddUser(t *testing.T) {
	user1, err := db.Transaction(func(session *xorm.Session) (i interface{}, e error) {
		user1 := Authuser{
			ID:          NewID().String(),
			Code:        "xxx",
			Name:        "xxx",
			Password:    "",
			Salt:        "",
			Avatar:      "",
			Orgid:       "",
			Email:       "",
			Phone:       "",
			Status:      "",
			Revision:    0,
			Createdby:   "admin",
			Createdtime: time.Now(),
			Updatedby:   "admin",
			Updatedtime: time.Now(),
		}
		if _, err := session.Insert(&user1); err != nil {
			return nil, err
		}
		user2 := Authuser{Name: "yyy"}
		if _, err := session.Where("ID = ?", user1.ID).Update(&user2); err != nil {
			return nil, err
		}
		user1.Name = user2.Name
		//if _, err := session.Exec("delete from AuthUser where Name = ?", user2.Name); err != nil {
		//	return nil, err
		//}
		return &user1, nil
	})

	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("TestAddUser: %v", user1)
	}
}
