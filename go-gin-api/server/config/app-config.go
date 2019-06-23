package config

import (
	"github.com/angenalZZZ/Go/go-gin-api/pkg/config"
	"github.com/angenalZZZ/Go/go-gin-api/pkg/config/app"
	"github.com/lexkong/log"
)

// App配置
var AppConfig *app.Config

// 初始化:App配置
func init() {
	if AppConfig == nil {
		AppConfig = new(app.Config)
		// 配置文件
		files := []string{"config.yml", "/etc/app/config.yml"}
		// 运行环境
		cfg := config.Config{Environ: &config.Environ{EnvironmentPrefix: "API"}}
		// 解析配置
		if err := cfg.Load(AppConfig, files...); err != nil {
			panic(err)
		}
		// 日志跟踪
		if err := log.InitWithConfig(&AppConfig.Log); err != nil {
			panic(err)
		}
	}
}
