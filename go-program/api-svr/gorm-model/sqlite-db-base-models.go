package gorm_model

import (
	"github.com/jinzhu/gorm"
)

type Foo struct {
	gorm.Model

	Bar int
	Baz string
}
