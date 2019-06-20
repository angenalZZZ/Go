package config

import (
	"github.com/angenalZZZ/Go/go-gin-api/pkg/config"
	"github.com/angenalZZZ/Go/go-gin-api/pkg/config/app"
)

// App 配置 变量
var appConfig *app.Config

// 获取 App 配置
func Get() *app.Config {
	if appConfig == nil {
		app, files := new(app.Config), []string{"config.yml", "/etc/app/config.yml"}
		cfg := config.Config{&config.Configuration{EnvironmentPrefix: "GI"}}
		err := cfg.Load(app, files...)
		if err != nil {
			panic(err)
		}
	}
	return appConfig
}
