package product

import (
	"fmt"
	"github.com/angenalZZZ/Go/go-gin-api/server/controller/param_bind"
	"github.com/angenalZZZ/Go/go-gin-api/server/controller/param_verify"
	"github.com/angenalZZZ/Go/go-gin-api/server/util/bind"
	"github.com/angenalZZZ/Go/go-gin-api/server/util/response"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

// 新增
func Add(c *gin.Context) {
	utilGin := response.Gin{Ctx: c}

	// 参数绑定
	s, e := bind.Bind(&param_bind.ProductAdd{}, c)
	if e != nil {
		utilGin.Response(-1, e.Error(), nil)
		return
	}

	// 参数验证
	validate := validator.New()

	// 注册自定义验证
	_ = validate.RegisterValidation("NameValid", param_verify.NameValid)

	if err := validate.Struct(s); err != nil {
		utilGin.Response(-1, err.Error(), nil)
		return
	}

	// 业务处理...

	utilGin.Response(1, "success", nil)
}

// 编辑
func Edit(c *gin.Context) {
	fmt.Println(c.Request.RequestURI)
}

// 删除
func Delete(c *gin.Context) {
	fmt.Println(c.Request.RequestURI)
}

// 详情

func Detail(c *gin.Context) {
	fmt.Println(c.Request.RequestURI)
}
