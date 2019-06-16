package config

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
)

type Config struct {
	*Configuration
}

type Configuration struct {
	Environment       string
	EnvironmentPrefix string
	// 异常：检查到错误配置
	ErrorOnUnmatchedKeys bool
}

// 解析：配置文件
func (c *Config) Load(config interface{}, files ...string) (err error) {
	configFiles := c.getConfigurationFiles(files...)

	for _, file := range configFiles {
		if err := DecodeFile(file, config, c.ErrorOnUnmatchedKeys); err != nil {
			return err
		}
	}

	prefix := c.EnvironmentPrefix
	if prefix == "" {
		return c.processTags(config)
	}
	return c.processTags(config, prefix)
}

// 异常：检查到错误配置
type UnmatchedTomlKeysError struct {
	Keys []toml.Key
}

func (e *UnmatchedTomlKeysError) Error() string {
	return fmt.Sprintf("检查到错误配置: %v", e.Keys)
}

func (c *Config) getConfigurationFiles(files ...string) []string {
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

func getFilenameWithEnvironmentPrefix(file, env string) (string, error) {
	extname := path.Ext(file)
	envFile := fmt.Sprintf("%v.%v%v", strings.TrimSuffix(file, extname), env, extname)

	if fileInfo, err := os.Stat(envFile); err == nil && fileInfo.Mode().IsRegular() {
		return envFile, nil
	}
	return "", errors.Errorf("加载配置文件失败：%v", file)
}
