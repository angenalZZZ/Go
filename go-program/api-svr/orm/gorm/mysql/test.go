package mysql

import (
	"github.com/angenalZZZ/Go/go-program/api-svr/orm/gorm/models"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

var ConnectionString string // https://github.com/go-sql-driver/mysql

func init() {
	ConnectionString = "root:HGJ766GR767FKJU0@/test?timeout=3s&charset=utf8&parseTime=True&loc=Local"
}

// 测试
func FooTestHandler(w http.ResponseWriter, r *http.Request) {
	models.FooTest(w, r, "mysql", ConnectionString)
}
