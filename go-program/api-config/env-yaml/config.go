package env_yaml

import (
	api_config "github.com/angenalZZZ/Go/go-program/api-config"
	"github.com/spf13/viper"
)

var config = &ApiConfigs{api_config.ApiConfigs{}}

type ApiConfigs struct {
	api_config.ApiConfigs
}

// 加载配置文件
func init() {
	// 设置全局变量
	api_config.Config = &config.ApiConfigs
	api_config.Check = config.Check

	viper.AutomaticEnv()
	viper.SetEnvPrefix("API")
}
