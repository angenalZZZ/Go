package config

import (
	"testing"

	"github.com/angenalZZZ/Go/go-gin-api/pkg/config/app"
	jsoniter "github.com/json-iterator/go"
)

func TestDecodeFile(t *testing.T) {
	config, file := new(app.Config), "app/config.example.yml"

	if err := DecodeFile(file, config, false); err != nil {
		t.Fatal(err)
	} else {
		t.Log(jsoniter.ConfigCompatibleWithStandardLibrary.MarshalToString(config))
	}
}
