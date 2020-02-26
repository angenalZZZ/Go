package product_dto

import "gopkg.in/go-playground/validator.v9"

// ProductAdd input
type ProductAdd struct {
	Name string `form:"name" json:"name" validate:"required,NameValid"`
}

// ProductAddValidate validator
func ProductAddValidate(input *ProductAdd) error {
	// 参数验证
	validate := validator.New()

	// 自定义验证
	_ = validate.RegisterValidation("NameValid", func(v validator.FieldLevel) bool {
		name := v.Field().String()
		if name == "admin" {
			return false
		}
		return true
	})

	return validate.Struct(input)
}
