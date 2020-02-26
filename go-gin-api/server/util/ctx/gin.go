package ctx

import (
	"github.com/angenalZZZ/Go/go-gin-api/server/model"
	"github.com/angenalZZZ/gofunc/http/errorcode"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// GinContext 请求上下文
type GinContext struct {
	C *gin.Context
}

// Wrap 获取请求上下文
func Wrap(c *gin.Context) *GinContext {
	return &GinContext{C: c}
}

// Bind 绑定输入参数
func (g *GinContext) Bind(input interface{}) error {
	b := binding.Default(g.C.Request.Method, g.C.ContentType())
	return g.C.ShouldBindWith(input, b)
}

// OK 处理成功
func (g *GinContext) OK(msg string, data interface{}) {
	g.C.JSON(errorcode.OK.GetHttpStatus(), model.Response{
		ErrorCode: *errorcode.OK.Msg(msg),
		Data:      data,
	})
	return
}

// Fail 处理失败
func (g *GinContext) Fail(msg string, data interface{}) {
	g.C.JSON(errorcode.INVALID.GetHttpStatus(), model.Response{
		ErrorCode: *errorcode.INVALID.Msg(msg),
		Data:      data,
	})
	return
}

// Error 处理时发生严重错误
func (g *GinContext) Error(msg string, data interface{}) {
	g.C.JSON(errorcode.ERROR.GetHttpStatus(), model.Response{
		ErrorCode: *errorcode.ERROR.Msg(msg),
		Data:      data,
	})
	return
}
