package gorm_model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"net/http"
	"os"
)

var sqliteDbPath string

func init() {
	var p = os.Getenv("GOPATH")
	sqliteDbPath = p + "/src/angenalZZZ/go-program/base.db"
	//println(sqliteDbPath)
}

// 测试
func SqliteTestBaseHandler(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", sqliteDbPath)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%v", err)))
		return
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Foo{})

	// Read
	var foo *Foo
	db.First(foo, "Bar=?", 1)

	// Insert
	if foo == nil {
		db.Create(&Foo{Bar: 1, Baz: "a"})
	}
	//db.First(foo, "Bar=?", 1)

	// Update
	//if foo != nil {
	//	db.Model(foo).Update("Baz", "b")
	//}

	// Delete
	//if foo != nil {
	//	db.Delete(foo)
	//}

	w.Write([]byte("successful to test database"))
}
