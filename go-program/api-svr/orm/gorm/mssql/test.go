package mssql

import (
	"angenalZZZ/go-program/api-svr/orm/gorm/models"

	//_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"net/http"
)

var ConnectionString string // https://github.com/denisenkom/go-mssqldb

func init() {
	ConnectionString = "sqlserver://sa:HGJ766GR767FKJU0@localhost/MSSQLSERVER?database=test&connection+timeout=3"
}

// 测试
func FooTestHandler(w http.ResponseWriter, r *http.Request) {
	models.FooTest(w, r, "mssql", ConnectionString)
}
