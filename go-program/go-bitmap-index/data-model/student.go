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
	if schema, err := db.Client.Schema(); err != nil {
		log.Panic(err)
	} else {
		Student = schema.Index("repository", pilosa.OptIndexKeys(true))
	}

	fields() // 属性|字段
}

// 初始化-属性|字段
func fields() {
	// 学时 时间量化：年月日时
	ST = Student.Field("ST", pilosa.OptFieldTypeTime(pilosa.TimeQuantumYearMonthDayHour))
}
