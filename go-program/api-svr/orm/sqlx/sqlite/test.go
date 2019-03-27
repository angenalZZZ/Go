package sqlite

import (
	"angenalZZZ/go-program/api-svr/orm/sqlx/models"

	"net/http"
)

var sqliteDbPath = ":memory:" // http://jmoiron.github.io/sqlx

func init() {
	//var p = os.Getenv("GOPATH")
	//sqliteDbPath = p + "/src/angenalZZZ/go-mock/test.db"
}

// 测试
func FooTestHandler(w http.ResponseWriter, r *http.Request) {
	models.FooTest(w, r, "sqlite3", sqliteDbPath)
}
