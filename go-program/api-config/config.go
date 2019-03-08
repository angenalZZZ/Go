package api_config

import (
	"github.com/gobuffalo/envy"
	"os"
)

type Config struct {
	HOST, POST, AUTH_JWT, JWT_algorithms, JWT_SECRET, REDISCLOUD_URL string
}

// 加载配置文件
func Load() {

	// *** 文件 .env 编码 必须是 UTF-8 +换行LF ***

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

	// 加载配置文件
	Load()

	// 检查配置项目
	check("HOST")
	check("POST")
	check("AUTH_JWT")
	check("JWT_algorithms")
	check("JWT_SECRET")
	check("REDISCLOUD_URL")
}

// 检查配置项
func check(key string) {
	if _, e := envy.MustGet(key); e != nil {
		panic(e)
	} else {
		//log.Printf("配置文件: %s = %s \n", key, val)
	}
}
