package config

import (
	"os"
	"testing"

	"github.com/angenalZZZ/Go/pkg/config/example"

	jsoniter "github.com/json-iterator/go"
)

func TestConfig_Load(t *testing.T) {
	// 设置环境变量
	err := os.Setenv("API_SERVER_ADDR", "api.server.com")
	if err != nil {
		t.Fatal(err)
	}
	err = os.Setenv("API_SERVER_PORT", "8080")
	if err != nil {
		t.Fatal(err)
	}

	// 读取配置信息
	config := Config{&Environ{EnvironmentPrefix: "API"}}
	appConfig, files := new(example.Config), []string{"example/config.example.yml"}

	// 解析配置文件
	err = config.Load(appConfig, files...)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log(jsoniter.ConfigCompatibleWithStandardLibrary.MarshalToString(appConfig))
	}
}
