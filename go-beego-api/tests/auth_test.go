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

// 3. 数据库连接客户端初始化
func init_auth_test() {
	db = conf.InitDbForXorm("mssql", "mssqlconn")
}

// 测试: 唯一标识生成器
func TestUUID(t *testing.T) {
	src1 := uuid.New()
	t.Logf("TestUUID: %s  %s", src1, pkg.NewIDFrom(src1))
}

// 测试: 保存用户信息到数据库
func TestAddUser(t *testing.T) {
	init_auth_test()
	users, err := db.Transaction(func(session *xorm.Session) (i interface{}, e error) {
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
		if _, e = session.Insert(&user1); e != nil {
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

	if err != nil {
		t.Fatal(err)
	} else if users != nil && len(users.([]auth.Authuser)) > 0 {
		t.Log(users)
	}
}
