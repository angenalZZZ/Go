package config

import (
	"io/ioutil"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/go-yaml/yaml"
	"github.com/json-iterator/go"
	"github.com/pkg/errors"
)

// DecodeFile 解析配置文件 转换为对象config
func DecodeFile(file string, config interface{}) (err error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	switch {
	case strings.HasSuffix(file, ".yml") || strings.HasSuffix(file, ".yaml"):
		return yaml.Unmarshal(data, config)

	case strings.HasSuffix(file, ".toml"):
		_, err := toml.Decode(string(data), config)
		return err

	case strings.HasSuffix(file, ".json"):
		return jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(data, config)

	default:
		return errors.Errorf("解析配置文件失败(*.yml *.toml *.json)：%v", file)
	}
}
