package conf

import (
	"errors"
	"github.com/astaxie/beego"
)

// 检查配置文件
func MustLoadAppConfig(path string) {
	if err := beego.LoadAppConfig("ini", path); err != nil {
		panic(err)
	}
}

// 获取配置文本
func GetAppConfig(name, path string) (i interface{}, e error) {
	i = beego.AppConfig.String(name)
	if i == "" {
		if e = beego.LoadAppConfig("ini", path); e != nil {
			return
		}
		i = beego.AppConfig.String(name)
	}
	if i == "" {
		e = errors.New("CONF \t\t" + name + " NOT FOUND")
	}
	return
}
