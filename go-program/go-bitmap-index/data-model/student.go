package data_model

import (
	db "github.com/angenalZZZ/Go/go-program/go-bitmap-index"
	log "github.com/sirupsen/logrus"

	"github.com/pilosa/go-pilosa"
)

// 学生
var (
	Student *pilosa.Index
	ST      *pilosa.Field
)

// 初始化
func init() {
	// 初始化-对象|数据
	schema, err := db.Client.Schema()
	if err != nil {
		log.Panic(err)
	} else {
		Student = schema.Index("repository", pilosa.OptIndexKeys(true))
	}

	// 初始化-属性|字段
	if err := fields(schema); err != nil {
		log.Panic(err)
	}
}

// 初始化-属性|字段
func fields(schema *pilosa.Schema) (err error) {
	// 学时 时间量化：年月日时
	ST = Student.Field("ST", pilosa.OptFieldTypeTime(pilosa.TimeQuantumYearMonthDayHour))

	ST.Set(0, 0)

	err = db.Client.SyncSchema(schema)
	return
}
