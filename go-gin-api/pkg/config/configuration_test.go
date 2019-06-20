package config

import (
	"testing"

	"github.com/angenalZZZ/Go/go-gin-api/pkg/config/app"

	jsoniter "github.com/json-iterator/go"
)

func TestConfig_Load(t *testing.T) {
	config := Config{&Configuration{EnvironmentPrefix: "GI"}}
	appConfig, files := new(app.Config), []string{"app/config.example.yml"}

	err := config.Load(appConfig, files...)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log(jsoniter.ConfigCompatibleWithStandardLibrary.MarshalToString(appConfig))
	}
}
