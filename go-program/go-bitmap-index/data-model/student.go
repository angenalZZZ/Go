package data_model

import (
	db "github.com/angenalZZZ/Go/go-program/go-bitmap-index"
	log "github.com/sirupsen/logrus"

	"github.com/pilosa/go-pilosa"
)

// 学生
var (
	/* 表格. Index → Sql数据库.Table */
	Student *pilosa.Index
	/* 字段. Field → Sql数据库.Column */
	ST      *pilosa.Field
	/* 行记录. Column  → Sql数据库.Row */
	/* 值..... Row  → Sql数据库.Value */
	/* 值&类型. Field.Value  → Sql数据库.Value(int) */
)

// 初始化 Table
func init() {
	// 初始化-对象|数据|Database/Table → 表格. Index
	schema, err := db.Client.Schema()
	if err != nil {
		log.Panic(err)
	} else {
		Student = schema.Index("student", pilosa.OptIndexKeys(true))
	}

	// 初始化-属性|字段|Table/Column → 字段. Field
	if err := fields(schema); err != nil {
		log.Panic(err)
	}
}

// 初始化 Table/Column
func fields(schema *pilosa.Schema) (err error) {
	// 学时 时间量化：年月日时
	ST = Student.Field("ST", pilosa.OptFieldTypeTime(pilosa.TimeQuantumYearMonthDayHour))

	ST.Set(0, 0)

	err = db.Client.SyncSchema(schema)
	return
}
