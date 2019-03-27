package models

import (
	"bytes"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"net/http"
)

/**
表结构 Foo
*/
type Foo struct {
	Bar int `db:"bar"` // 首字母必须大写
	Baz sql.NullString
}

// 测试
func FooTest(w http.ResponseWriter, r *http.Request, dbType, connStr string) {
	db, err := sqlx.Open(dbType, connStr)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%v", err)))
		return
	}
	defer db.Close()

	var buf bytes.Buffer

	// Migrate the schema
	//db.AutoMigrate(&Foo{})
	//buf.WriteString(" AutoMigrate\n")
	db.MustExec(`
CREATE TABLE IF NOT EXISTS foo (
    bar int not null,
    baz text
);`)

	// Delete rows
	if i, _ := db.MustExec(`delete from foo`).RowsAffected(); i > 0 {
		buf.WriteString(fmt.Sprintf(" Remove rows:%d\n", i))
	}

	// Read
	foo := Foo{}
	foos := []Foo{}
	if e := db.Get(&foo, `SELECT * FROM foo WHERE bar>? LIMIT 1`, foo.Bar); e != sql.ErrNoRows || foo.Bar > 0 {
		buf.WriteString(fmt.Sprintf(" First(bar=1):%+v\n%v\n", foo, e))
	}

	// Create insert
	if foo.Bar == 0 {
		foo = Foo{Bar: 1, Baz: sql.NullString{String: "a", Valid: true}}
		buf.WriteString(fmt.Sprintf(" Insert:%+v\n", foo))
		if i, e := db.MustExec(`INSERT INTO foo VALUES (?, ?)`, foo.Bar, foo.Baz).RowsAffected(); i > 0 {
			buf.WriteString(" Inserted Ok\n")
		} else {
			buf.WriteString(fmt.Sprintf(" Not Inserted\n%v\n", e))
		}
	}
	if e := db.Select(&foos, `SELECT * FROM foo WHERE bar=? and baz=?`, foo.Bar, foo.Baz); len(foos) > 0 {
		foo = foos[0]
		buf.WriteString(fmt.Sprintf(" Inserted:Find %d rows\n", len(foos)))
	} else {
		foo = Foo{}
		buf.WriteString(fmt.Sprintf(" Inserted:Not Find\n%v\n", e))
	}

	// Update
	if foo.Bar > 0 {
		if i, e := db.MustExec(`UPDATE foo SET baz=? WHERE bar=?`, "b", 1).RowsAffected(); i > 0 {
			buf.WriteString(" Updated Ok\n")
		} else {
			buf.WriteString(fmt.Sprintf(" Not Updated\n%v\n", e))
		}
	}

	// Delete
	if foo.Bar > 0 {
		if i, e := db.MustExec(`DELETE FROM foo WHERE bar=?`, foo.Bar).RowsAffected(); i > 0 {
			buf.WriteString(" Deleted Ok\n")
		} else {
			buf.WriteString(fmt.Sprintf(" Not Deleted\n%v\n", e))
		}
	}

	// Response
	w.Write(buf.Bytes())
}
