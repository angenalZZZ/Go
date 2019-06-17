package config

import (
	"github.com/json-iterator/go"
	"testing"
)

func TestDecodeFile(t *testing.T) {
	config, file := new(AppConfiguration), "../config.example.yml"

	if err := DecodeFile(file, config, false); err != nil {
		t.Fatal(err)
	} else {
		t.Log(jsoniter.ConfigCompatibleWithStandardLibrary.MarshalToString(config))
	}
}
