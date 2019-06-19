package config

import (
	"github.com/json-iterator/go"
	"testing"
)

func TestConfig_Load(t *testing.T) {
	config := Config{&Configuration{EnvironmentPrefix: "GI"}}
	appConfig, files := new(AppConfiguration), []string{"config.example.yml"}

	err := config.Load(appConfig, files...)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log(jsoniter.ConfigCompatibleWithStandardLibrary.MarshalToString(appConfig))
	}
}
