package conf

import "github.com/astaxie/beego"

// 检查配置文件
func MustLoadAppConfig(path string) {
	if err := beego.LoadAppConfig("ini", path); err != nil {
		panic(err)
	}
}
