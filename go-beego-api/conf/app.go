package conf

import (
	"errors"
	"github.com/astaxie/beego"
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
