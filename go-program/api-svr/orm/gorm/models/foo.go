package models

import (
	"bytes"
	"fmt"
	"github.com/jinzhu/gorm"
	"net/http"
)

/**
表结构 Foo
*/
type Foo struct {
	gorm.Model

	Bar int
	Baz string
}

// 测试
func FooTest(w http.ResponseWriter, r *http.Request, dbType, connStr string) {
	db, err := gorm.Open(dbType, connStr)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%v", err)))
		return
	}
	defer db.Close()

	var buf bytes.Buffer

	// Migrate the schema
	db.AutoMigrate(&Foo{})
	buf.WriteString(" AutoMigrate\n")

	// Delete, not Remove rows
	if i := db.Table("foos").Delete(Foo{}, "1=1").RowsAffected; i > 0 {
		buf.WriteString(fmt.Sprintf(" Delete rows:%d\n", i))
	}
	// Delete, and Remove rows
	if i := db.Exec("delete from foos").RowsAffected; i > 0 {
		buf.WriteString(fmt.Sprintf(" Remove rows:%d\n", i))
	}

	// Read
	var foo *Foo
	var foos []Foo
	if i := db.Where("bar=?", 1).First(foo).RowsAffected; i > 0 {
		buf.WriteString(fmt.Sprintf(" First(bar=1):%+v\n", foo))
	}

	// Create insert
	if foo == nil {
		foo1 := &Foo{Bar: 1, Baz: "a"}
		buf.WriteString(fmt.Sprintf(" Insert:%+v\n", struct {
			Bar int
			Baz string
		}{foo1.Bar, foo1.Baz}))
		if i := db.Create(foo1).RowsAffected; i > 0 {
			buf.WriteString(" Inserted Ok\n")
		} else {
			buf.WriteString(" Not Inserted\n")
		}
	}
	if i := db.Where("bar=? and baz=?", 1, "a").Find(&foos).RowsAffected; i > 0 {
		foo = &foos[0]
		buf.WriteString(fmt.Sprintf(" Inserted:Find %d rows\n", i))
	} else {
		foo = nil
		buf.WriteString(" Inserted:Not Find\n")
	}

	// Update
	if foo != nil {
		if i := db.Model(foo).Update("baz", "b").RowsAffected; i > 0 {
			buf.WriteString(" Updated Ok\n")
		} else {
			buf.WriteString(" Not Updated\n")
		}
	}

	// Delete
	if foo != nil {
		if i := db.Delete(foo).RowsAffected; i > 0 {
			buf.WriteString(" Deleted Ok\n")
		} else {
			buf.WriteString(" Not Deleted\n")
		}
	}

	// Response
	w.Write(buf.Bytes())
}
