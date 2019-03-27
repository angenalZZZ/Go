package sqlite

import (
	"angenalZZZ/go-program/api-svr/orm/gorm/models"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"net/http"
)

var sqliteDbPath = ":memory:" // https://github.com/mattn/go-sqlite3

func init() {
	//var p = os.Getenv("GOPATH")
	//sqliteDbPath = p + "/src/angenalZZZ/go-mock/test.db"
}

// 测试
func FooTestHandler(w http.ResponseWriter, r *http.Request) {
	models.FooTest(w, r, "sqlite3", sqliteDbPath)
}
