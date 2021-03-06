package config

import (
	"github.com/angenalZZZ/gofunc/configfile"
	"github.com/angenalZZZ/gofunc/f"
	"github.com/angenalZZZ/gofunc/log"
	"github.com/spf13/pflag"
)

// App配置
var (
	AppConfig     *AppConfigModel
	AppConfigFile = pflag.StringP("config", "c", "server.config.yml", "api server config file path.")

	ApiAuthConfig = map[string]map[string]string{

		// 调用方
		"DEMO": {
			"md5": "IgkibX71IEf382PT",
			"aes": "IgkibX71IEf382PT",
			"rsa": "rsa/public.pem",
		},
	}

	// 签名超时时间
	AppSignExpiry = "120"

	// RSA Private File
	AppRsaPrivateFile = "rsa/private.pem"
)

// init 初始化:App配置
func init() {
	if AppConfig == nil {
		AppConfig = new(AppConfigModel)
		// 解析配置
		f.Must(configfile.YamlTo(*AppConfigFile, AppConfig))
		// 日志跟踪
		f.Must(log.InitWithConfig(&AppConfig.Log))
	}
}
