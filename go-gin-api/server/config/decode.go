package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/go-yaml/yaml"
	"github.com/json-iterator/go"
	"github.com/pkg/errors"
)

// 解析配置文件 转换为config
func DecodeFile(file string, config interface{}, errorOnUnmatchedKeys bool) (err error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	switch {
	case strings.HasSuffix(file, ".yml") || strings.HasSuffix(file, ".yaml"):
		if errorOnUnmatchedKeys {
			return yaml.UnmarshalStrict(data, config)
		}
		return yaml.Unmarshal(data, config)

	case strings.HasSuffix(file, ".toml"):
		if metadata, err := toml.Decode(string(data), config); err == nil && len(metadata.Undecoded()) > 0 && errorOnUnmatchedKeys {
			return &UnmatchedTomlKeysError{Keys: metadata.Undecoded()}
		}
		return err

	case strings.HasSuffix(file, ".json"):
		return jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(data, config)

	default:
		return errors.Errorf("解析配置文件失败：%v", file)
	}
}

// 解析配置结构 处理config标记
func (c *Config) processTags(config interface{}, prefixes ...string) (err error) {
	configValue := reflect.Indirect(reflect.ValueOf(config))

	// 当配置项为指针时,读取其结构
	for configValue.Kind() == reflect.Ptr {
		configValue = configValue.Elem()
	}

	// 当配置项非结构时,抛出异常
	if configValue.Kind() != reflect.Struct {
		return errors.Errorf("解析配置时,读取其结构发生错误：%v", configValue.Kind().String())
	}

	// 解析配置项数据结构
	configType := configValue.Type()
	for i := 0; i < configType.NumField(); i++ {
		var (
			structField = configType.Field(i)
			valueField  = configValue.Field(i)
			envName     = structField.Tag.Get("env")
			envNames    []string
		)

		if !valueField.CanAddr() || !valueField.CanInterface() {
			continue
		}

		if envName == "" {
			envNames = append(envNames, strings.Join(append(prefixes, structField.Name), "_"))
			envNames = append(envNames, strings.ToUpper(strings.Join(append(prefixes, structField.Name), "_")))
		} else {
			envNames = []string{envName}
		}

		for _, env := range envNames {
			if value := os.Getenv(env); value != "" {
				if err = yaml.Unmarshal([]byte(value), valueField.Addr().Interface()); err != nil {
					return err
				}
				break
			}
		}

		// 配置项为空时
		if isZero := reflect.DeepEqual(valueField.Interface(), reflect.Zero(valueField.Type()).Interface()); isZero {
			// 配置项存在默认值时 标记：`default:""`
			if value := structField.Tag.Get("default"); value != "" {
				if err = yaml.Unmarshal([]byte(value), valueField.Addr().Interface()); err != nil {
					return err
				}
			} else if structField.Tag.Get("required") == "true" {
				return errors.Errorf("解析配置时缺少必填项：%v", structField.Name)
			}
		}

		for valueField.Kind() == reflect.Ptr {
			valueField = valueField.Elem()
		}

		if valueField.Kind() == reflect.Struct {
			err = c.processTags(
				valueField.Addr().Interface(),
				prefix(prefixes, &structField)...,
			)
			if err != nil {
				return err
			}
		}

		if valueField.Kind() == reflect.Slice {
			for i := 0; i < valueField.Len(); i++ {
				if reflect.Indirect(valueField.Index(i)).Kind() == reflect.Struct {
					err = c.processTags(
						valueField.Index(i).Addr().Interface(),
						append(prefix(prefixes, &structField), fmt.Sprint(i))...,
					)
					if err != nil {
						return err
					}
				}
			}
		}
	}
	return nil
}

func prefix(prefixes []string, structField *reflect.StructField) []string {
	if structField.Anonymous && structField.Tag.Get("anonymous") == "true" {
		return prefixes
	}
	return append(prefixes, structField.Name)
}
