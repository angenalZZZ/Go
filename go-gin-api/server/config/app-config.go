package config

import (
	"github.com/angenalZZZ/Go/go-gin-api/pkg/config"
	"github.com/angenalZZZ/Go/go-gin-api/pkg/config/app"
	"github.com/lexkong/log"
	"github.com/spf13/pflag"
)

// App配置
var (
	AppConfig     *app.Config
	AppConfigFile = pflag.StringP("config", "c", "config.yml", "api server config file path.")
)

// 初始化:App配置
func init() {
	if AppConfig == nil {
		AppConfig = new(app.Config)
		// 环境配置
		cfg := config.Config{Environ: &config.Environ{EnvironmentPrefix: "API"}}
		// 解析配置
		if err := cfg.Load(AppConfig, *AppConfigFile); err != nil {
			panic(err)
		}
		// 日志跟踪
		if err := log.InitWithConfig(&AppConfig.Log); err != nil {
			panic(err)
		}
	}
}
