package config

import (
	"testing"

	"github.com/angenalZZZ/Go/pkg/config/example"
	jsoniter "github.com/json-iterator/go"
)

func TestDecodeFile(t *testing.T) {
	config, file := new(example.Config), "example/config.example.yml"

	if err := DecodeFile(file, config); err != nil {
		t.Fatal(err)
	} else {
		t.Log(jsoniter.ConfigCompatibleWithStandardLibrary.MarshalToString(config))
	}
}
