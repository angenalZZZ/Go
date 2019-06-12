package mode

import "github.com/gin-gonic/gin"

const (
	// 开发调试模式
	Debug = "debug"
	// 测试检测模式
	Test = "test"
	// 发布生产模式
	Prod = "prod"
)

var mode = Debug

// 当前环境设置
func Set(newMode string) {
	mode = newMode
	updateGinMode()
}

// 当前环境获取
func Get() string {
	return mode
}

// 当前为开发或测试环境
func IsDev() bool {
	return Get() == Debug || Get() == Test
}

func updateGinMode() {
	switch Get() {
	case Debug:
		gin.SetMode(gin.DebugMode)
	case Test:
		gin.SetMode(gin.TestMode)
	case Prod:
		gin.SetMode(gin.ReleaseMode)
	default:
		panic("当前环境设置异常")
	}
}
