package route

import (
	"github.com/angenalZZZ/Go/go-gin-api/server/controller/jaeger_conn"
	"github.com/angenalZZZ/Go/go-gin-api/server/controller/product"
	"github.com/angenalZZZ/Go/go-gin-api/server/controller/test"
	"github.com/angenalZZZ/Go/go-gin-api/server/route/middleware/exception"
	"github.com/angenalZZZ/Go/go-gin-api/server/route/middleware/jaeger"
	"github.com/angenalZZZ/Go/go-gin-api/server/route/middleware/logger"
	"github.com/angenalZZZ/Go/go-gin-api/server/util/response"
	"github.com/gin-gonic/gin"
)

func SetupRouter(engine *gin.Engine) {

	//设置路由中间件
	engine.Use(logger.SetUp(), exception.SetUp(), jaeger.SetUp())

	//404
	engine.NoRoute(func(c *gin.Context) {
		utilGin := response.Gin{Ctx: c}
		utilGin.Response(404, "请求方法不存在", nil)
	})

	engine.GET("/ping", func(c *gin.Context) {
		utilGin := response.Gin{Ctx: c}
		utilGin.Response(1, "pong", nil)
	})

	// 测试链路追踪
	engine.GET("/jaeger_test", jaeger_conn.JaegerTest)

	//@todo 记录请求超时的路由

	ProductRouter := engine.Group("/product")
	{
		// 新增产品
		ProductRouter.POST("", product.Add)

		// 更新产品
		ProductRouter.PUT("/:id", product.Edit)

		// 删除产品
		ProductRouter.DELETE("/:id", product.Delete)

		// 获取产品详情
		ProductRouter.GET("/:id", product.Detail)
	}

	// 测试加密性能
	TestRouter := engine.Group("/test")
	{
		// 测试 MD5 组合 的性能
		TestRouter.GET("/md5", test.Md5Test)

		// 测试 AES 的性能
		TestRouter.GET("/aes", test.AesTest)

		// 测试 RSA 的性能
		TestRouter.GET("/rsa", test.RsaTest)
	}
}
