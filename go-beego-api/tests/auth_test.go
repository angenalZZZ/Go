package test

import (
	"github.com/angenalZZZ/Go/go-beego-api/conf"
	"github.com/angenalZZZ/Go/go-beego-api/models/auth"
	"github.com/angenalZZZ/Go/go-beego-api/pkg"
	"github.com/google/uuid"
	"github.com/xormplus/xorm"
	"testing"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
)

// 1. 生成数据库实体models - DbFirst 工具 - https://github.com/go-xorm/cmd/xorm
//    > cp %GOPATH%/src/github.com/go-xorm/cmd/xorm/templates/goxorm/* ./_templates/goxorm
//    > xorm reverse mssql "server=localhost;user id=sa;password=HGJ766GR767FKJU0;database=AppAuth" ./_templates/goxorm ./models/auth ^Auth

// 2. 数据库引擎增强版 db orm engine
var db *xorm.Engine

const (
	mssqlDriverName = "mssql"
	mssqlConn       = "mssqlconn"
)

// 3. 数据库连接客户端初始化
func initAuthTest() {
	db = conf.InitDbForXOrm(mssqlDriverName, mssqlConn)
}

func initUser1() *auth.Authuser {
	user1 := auth.Authuser{
		Id:          pkg.NewID().String(),
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
	return &user1
}

// 测试: 唯一标识生成器
func TestUUID(t *testing.T) {
	src1 := uuid.New()
	t.Logf("TestUUID: %s  %s", src1, pkg.NewIDFrom(src1))
}

// 测试: 保存用户信息
func TestAuthuser(t *testing.T) {
	initAuthTest()
	db := auth.AuthuserXorm{DB: db}

	user1 := initUser1()
	if e := db.Create(user1); e != nil {
		t.Fatal(e)
		return
	}
	user2, e := db.GetById(user1.Id)
	if e != nil {
		t.Fatal(e)
		return
	}
	user2.Email = "yyy@qq.com"
	//if e := db.Update(user2, "Email"); e != nil {
	//	t.Fatal(e)
	//	return
	//}
	//if _, e := db.GetByEmailOrPhone(user2.Email); e != nil {
	//	t.Fatal(e)
	//	return
	//}
	//if e := db.Delete(user2.Id); e != nil {
	//	t.Fatal(e)
	//	return
	//}
}

// 测试: 保存用户信息的事务处理
func TestAuthuserTransaction(t *testing.T) {
	initAuthTest()

	users, e := db.Transaction(func(session *xorm.Session) (i interface{}, e error) {
		user1 := initUser1()
		if _, e = session.Insert(user1); e != nil {
			return
		}
		user2 := auth.Authuser{Name: "yyy"}
		if _, e = session.Where("Id=?", user1.Id).Update(&user2); e != nil {
			return
		}
		if _, e = session.ID(user1.Id).Get(&user2); e != nil {
			return
		}
		if _, e = session.Exec("delete from AuthUser where Id=?", user2.Id); e != nil {
			return
		}
		users := make([]auth.Authuser, 0, 1)
		if e = session.Cols("Id", "Code", "Name").Or("Name=?", user1.Name).Or("Name=?", user2.Name).Limit(10, 0).Find(&users); e != nil {
			return
		}
		if e = pkg.NewPagingBuilder(func(builder *pkg.Pager) {
			builder.Select("Id", "Code", "Name").From("AuthUser").OrderBy("Name ASC")
		}).Paging(session, 10, 10, &users); e != nil {
			return
		}
		i = users
		return
	})

	if e != nil {
		t.Fatal(e)
	} else if users != nil && len(users.([]auth.Authuser)) > 0 {
		t.Log(users)
	}
}
