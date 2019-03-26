package mssql

import (
	"angenalZZZ/go-program/api-svr/gorm/models"

	_ "github.com/jinzhu/gorm/dialects/mssql"
	"net/http"
)

var ConnectionString string

func init() {
	ConnectionString = "sqlserver://sa:123456@localhost:1433?database=test"
}

// 测试
func FooTestHandler(w http.ResponseWriter, r *http.Request) {
	models.FooTest(w, r, "mssql", ConnectionString)
}
