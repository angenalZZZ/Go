package config

import (
	"github.com/angenalZZZ/Go/go-gin-api/pkg/config"
	"github.com/angenalZZZ/Go/go-gin-api/pkg/config/app"
)

// App配置
var AppConfig *app.Config

// 初始化:App配置
func init() {
	if AppConfig == nil {
		AppConfig = new(app.Config)
		files := []string{"config.yml", "/etc/app/config.yml"}
		cfg := config.Config{Environ: &config.Environ{EnvironmentPrefix: "API"}}
		err := cfg.Load(AppConfig, files...)
		if err != nil {
			panic(err)
		}
	}
}
