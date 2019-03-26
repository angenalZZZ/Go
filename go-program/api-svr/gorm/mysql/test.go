package mysql

import (
	"angenalZZZ/go-program/api-svr/gorm/models"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

var ConnectionString string

func init() {
	ConnectionString = "root:123456@/localhost:3306/test?timeout=3s&charset=utf8&parseTime=True&loc=Local"
}

// 测试
func FooTestHandler(w http.ResponseWriter, r *http.Request) {
	models.FooTest(w, r, "mysql", ConnectionString)
}
