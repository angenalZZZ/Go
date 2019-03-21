package api_config

import (
	//"github.com/chai2010/winsvc"
	"github.com/gobuffalo/envy"
	"log"
	"os"
	"strconv"
)

var JwtConf *JwtConfig

// 认证方式 Jwt Config: 复合类型
type JwtConfig struct {
	AUTH_JWT, JWT_algorithms, JWT_SECRET string
	JWT_LIFETIME                         int
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
	if JwtConf == nil {
		println()

		// 检查配置项目
		Check("AUTH_JWT")
		Check("JWT_algorithms")
		Check("JWT_SECRET")
		Check("JWT_LIFETIME")
		lifetime, _ := strconv.Atoi(os.Getenv("JWT_LIFETIME"))
		JwtConf = &JwtConfig{
			AUTH_JWT:       os.Getenv("AUTH_JWT"),
			JWT_algorithms: os.Getenv("JWT_algorithms"),
			JWT_SECRET:     os.Getenv("JWT_SECRET"),
			JWT_LIFETIME:   lifetime,
		}
		log.Printf("加载配置文件并检查配置项: OK\n")
	}
}

// 检查配置项 Must Checked
func Check(key string) {
	if _, e := envy.MustGet(key); e != nil {
		panic(e)
	} else {
		//log.Printf("配置文件: %s = %s \n", key, val)
	}
}
