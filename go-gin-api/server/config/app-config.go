package config

import (
	"github.com/angenalZZZ/Go/pkg"
	"github.com/angenalZZZ/Go/pkg/config"
	"github.com/lexkong/log"
	"github.com/spf13/pflag"
)

// App配置
var (
	AppConfig     *Config
	AppConfigFile = pflag.StringP("config", "c", "app.config.yml", "api server config file path.")
)

// init 初始化:App配置
func init() {
	if AppConfig == nil {
		AppConfig = new(Config)
		// 环境配置
		cfg := config.Config{Environ: &config.Environ{EnvironmentPrefix: "API"}}
		// 解析配置
		pkg.MustNotError(cfg.Load(AppConfig, *AppConfigFile))
		// 日志跟踪
		pkg.MustNotError(log.InitWithConfig(&AppConfig.Log))
	}
}
