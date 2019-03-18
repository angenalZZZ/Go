package api_config

import (
	//"github.com/chai2010/winsvc"
	"github.com/gobuffalo/envy"
	"log"
	"os"
)

var Env Config

// 认证方式 结构设计 [Config:复合类型]
type Config struct {
	AUTH_JWT, JWT_algorithms, JWT_SECRET string
}

// 加载配置文件
func init() {

	// *** 文件 .env 编码 必须是 UTF-8 +换行LF ***
	//AppPath, _ = winsvc.GetAppPath()

	// 检查环境，判断加载的配置文件
	f, v := ".env.prod", os.Getenv("GO_ENV")
	if v == "development" {
		f = ".env"
	}
	//log.Printf("配置文件: 解析 %s\n", f)

	// 配置文件错误时，直接退出应用
	if e := envy.Load(f); e != nil {
		panic(e)
	} else if v == "development" {
		//for _, v := range envy.Environ() {
		//	println(v)
		//}
	}
}

// 加载配置文件并检查配置项
func LoadCheck() {
	println()

	// 检查配置项目
	Check("AUTH_JWT")
	Check("JWT_algorithms")
	Check("JWT_SECRET")
	Env = Config{
		AUTH_JWT:       os.Getenv("AUTH_JWT"),
		JWT_algorithms: os.Getenv("JWT_algorithms"),
		JWT_SECRET:     os.Getenv("JWT_SECRET"),
	}
	log.Printf("加载配置文件并检查配置项: OK\n")
}

// 检查配置项 Must Checked
func Check(key string) {
	if _, e := envy.MustGet(key); e != nil {
		panic(e)
	} else {
		//log.Printf("配置文件: %s = %s \n", key, val)
	}
}
