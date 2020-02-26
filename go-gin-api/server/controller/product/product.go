package product

import (
	dto "github.com/angenalZZZ/Go/go-gin-api/server/controller/product_dto"
	"github.com/angenalZZZ/Go/go-gin-api/server/util/ctx"
	"github.com/gin-gonic/gin"
)

// 新增
func Add(c *gin.Context) {
	g := ctx.Wrap(c)

	// 参数绑定
	s := new(dto.ProductAdd)
	if e := g.Bind(s); e != nil {
		g.Fail(e.Error(), nil)
		return
	}

	// 参数验证
	if err := dto.ProductAddValidate(s); err != nil {
		g.Fail(err.Error(), nil)
		return
	}

	// 业务处理...

	g.OK("success", nil)
}

// 编辑
func Edit(c *gin.Context) {
	g := ctx.Wrap(c)

	g.OK(c.Request.RequestURI, nil)
}

// 删除
func Delete(c *gin.Context) {
	g := ctx.Wrap(c)

	g.OK(c.Request.RequestURI, nil)
}

// 详情
func Detail(c *gin.Context) {
	g := ctx.Wrap(c)

	g.OK(c.Request.RequestURI, nil)
}
