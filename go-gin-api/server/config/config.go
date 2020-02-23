package config

import (
	"fmt"
	"os"
	"path"
	"reflect"
	"strings"

	"github.com/go-yaml/yaml"
	"github.com/pkg/errors"
)

// Config 配置功能
type Config struct {
	// 运行环境
	*Environ
}

// Environ 配置运行环境
type Environ struct {
	// 当前环境
	Environment string
	// 环境变量：前缀
	EnvironmentPrefix string
}

// Load 解析配置文件
func (c *Config) Load(config interface{}, files ...string) (err error) {
	// 查找配置文件
	configFiles := c.GetConfigFiles(files...)

	// 解析配置文件 转换为对象config
	for _, file := range configFiles {
		if err := DecodeFile(file, config); err != nil {
			return err
		}
	}

	// 解析配置文件数据结构：标记
	return c.DecodeTags(config)
}

// GetConfigFiles 查找配置文件
func (c *Config) GetConfigFiles(files ...string) []string {
	var results []string

	for i := len(files) - 1; i >= 0; i-- {
		file, found := files[i], false

		// 检查[默认]配置文件 : config.yml
		if fileInfo, err := os.Stat(file); err == nil && fileInfo.Mode().IsRegular() {
			found = true
			results = append(results, file)
		}

		// 检查[环境]配置文件 : config.*.yml
		if fileWithPrefix, err := getFilenameWithEnvironmentPrefix(file, c.Environment); err == nil {
			found = true
			results = append(results, fileWithPrefix)
		}

		// 检查[例子]配置文件 : config.example.yml
		if !found {
			if fileWithPrefix, err := getFilenameWithEnvironmentPrefix(file, "example"); err == nil {
				results = append(results, fileWithPrefix)
			}
		}
	}
	return results
}

// DecodeTags 解析配置文件数据结构：标记
func (c *Config) DecodeTags(config interface{}, prefixes ...string) (err error) {
	if len(prefixes) == 0 && c.EnvironmentPrefix != "" {
		prefixes = []string{c.EnvironmentPrefix} // 环境变量：前缀
	}

	configValue := reflect.Indirect(reflect.ValueOf(config))

	// 当配置项为指针时,读取数据结构
	for configValue.Kind() == reflect.Ptr {
		configValue = configValue.Elem()
	}

	// 当配置项非结构时,抛出异常
	if configValue.Kind() != reflect.Struct {
		return errors.Errorf("解析配置时异常,无法读取数据结构：%v", configValue.Kind().String())
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

		// 检查环境变量
		if envName == "" {
			// 按数据结构路径名称
			envNames = append(envNames, strings.Join(append(prefixes, structField.Name), "_"))
			envNames = append(envNames, strings.ToUpper(strings.Join(append(prefixes, structField.Name), "_")))
		} else {
			// 按指定名称：`env:""`
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
			if value := structField.Tag.Get("default"); value != "" {
				// 有默认值时：`default:""`
				if err = yaml.Unmarshal([]byte(value), valueField.Addr().Interface()); err != nil {
					return err
				}
			} else if structField.Tag.Get("required") == "true" {
				// 为必填项时：`required:"true"`
				return errors.Errorf("解析配置时异常,缺少必填项：%v (%v)", structField.Name, structField.Type)
			}
		}

		for valueField.Kind() == reflect.Ptr {
			valueField = valueField.Elem()
		}

		if valueField.Kind() == reflect.Struct {
			err = c.DecodeTags(
				valueField.Addr().Interface(),
				prefixToDecodeTags(prefixes, &structField)...,
			)
			if err != nil {
				return err
			}
		}

		if valueField.Kind() == reflect.Slice {
			for i := 0; i < valueField.Len(); i++ {
				if reflect.Indirect(valueField.Index(i)).Kind() == reflect.Struct {
					err = c.DecodeTags(
						valueField.Index(i).Addr().Interface(),
						append(prefixToDecodeTags(prefixes, &structField), fmt.Sprint(i))...,
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

// 查找特定环境配置文件
func getFilenameWithEnvironmentPrefix(file, env string) (string, error) {
	extname := path.Ext(file)
	envFile := fmt.Sprintf("%v.%v%v", strings.TrimSuffix(file, extname), env, extname)

	if fileInfo, err := os.Stat(envFile); err == nil && fileInfo.Mode().IsRegular() {
		return envFile, nil
	}
	return "", errors.Errorf("未找到特定环境[%s]配置文件：%s", env, file)
}

func prefixToDecodeTags(prefixes []string, structField *reflect.StructField) []string {
	if structField.Anonymous && structField.Tag.Get("anonymous") == "true" {
		return prefixes
	}
	return append(prefixes, structField.Name)
}
