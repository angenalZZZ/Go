package conf

import (
	"errors"
	"github.com/astaxie/beego"
	"path/filepath"
	"runtime"
)

// 获取配置信息 from beego.AppConfig
func GetAppConfig(name, path string) (i interface{}, e error) {
	i = beego.AppConfig.String(name)
	if i == "" {
		if e = beego.LoadAppConfig("ini", path); e != nil {
			return
		}
		i = beego.AppConfig.String(name)
	}
	if i == "" {
		e = errors.New("beego.AppConfig \t\t" + name + " NOT FOUND")
	}
	return
}

// 获取工作目录
func GetAppPath() string {
	_, file, _, _ := runtime.Caller(0)
	path, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	return path
}

// 获取配置文件路径
func GetAppConfigPath(conf string) string {
	return filepath.Join(GetAppPath(), "conf", conf)
}
