package env_envy

import (
	"net"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/seefan/gossdb/conf"

	api_config "github.com/angenalZZZ/Go/go-program/api-config"

	"os"
	"strconv"

	"github.com/gobuffalo/envy"
)

var config = &ApiConfigs{} // 当前配置 访问对象

type ApiConfigs struct {
	api_config.ApiConfigs  // 当前配置 所有变量
	api_config.IApiConfigs // 当前配置 需实现的方法
}

// 加载配置文件 插件 github.com/gobuffalo/envy
func init() {
	// 加载配置文件并检查配置项
	config.Load()

	// 当前配置 设置静态变量
	api_config.Config = &config.ApiConfigs
}

// 加载配置文件并检查配置项
func (c *ApiConfigs) Load() {
	// *** 文件 .env 编码 必须是 UTF-8 +换行LF ***
	//AppPath, _ = winsvc.GetAppPath()

	// 检查环境，判断加载的配置文件 (GO_ENV:环境变量或命令行参数)
	f, v := ".env.prod", os.Getenv("GO_ENV")
	if v == "development" {
		f = ".env.dev"
	} else if v == "test" {
		f = ".env.test"
	}
	//log.Printf("配置文件: 解析 %s\n", f)

	// 配置文件错误时，直接退出应用
	if e := envy.Load(f); e != nil {
		panic(e)
	}

	println()

	// 检查配置项目
	if c.Jwt == nil {
		c.Check("AUTH_JWT")
		c.Check("JWT_algorithms")
		c.Check("JWT_SECRET")
		c.Check("JWT_LIFETIME")
		c.Check("JWT_Issuer")
		c.Check("JWT_Subject")
		c.Check("JWT_Audience")
		c.Check("JWT_Sign")

		lifetime, _ := strconv.Atoi(os.Getenv("JWT_LIFETIME"))

		jwtSign := api_config.JwtSign{}
		s := strings.Split(os.Getenv("JWT_Sign"), ",")
		if len(s) == 3 {
			jwtSign.Key, jwtSign.Pub, jwtSign.Alg = s[0], s[1], s[2]
		}

		c.Jwt = &api_config.JwtConfig{
			AUTH_JWT:       os.Getenv("AUTH_JWT"),
			JWT_algorithms: os.Getenv("JWT_algorithms"),
			JWT_SECRET:     os.Getenv("JWT_SECRET"),
			JWT_LIFETIME:   lifetime,
			JWT_Issuer:     os.Getenv("JWT_Issuer"),
			JWT_Subject:    os.Getenv("JWT_Subject"),
			JWT_Audience:   os.Getenv("JWT_Audience"),
			JWT_Sign:       jwtSign,
		}
	}

	// HttpWeb服务配置
	if c.HttpSvr == nil {
		c.Check("HOST")
		c.Check("POST")

		c.HttpSvr = &api_config.HttpSvrConfig{
			Addr: os.Getenv("HOST") + ":" + os.Getenv("POST"),
		}
	}

	// Redis 客户端接口
	if c.RedisCli == nil {

		db, e := strconv.Atoi(os.Getenv("REDIS_DB"))
		if e != nil {
			db = 0
		}

		c.RedisCli = &api_config.RedisCliConfig{
			Addr: os.Getenv("REDIS_ADDR"),
			Pwd:  os.Getenv("REDIS_PWD"),
			Db:   db,
		}
	}

	// LevelDb 存储目录地址
	if c.LevelDb == nil {
		c.Check("LEVELDB")

		c.LevelDb = &api_config.LevelDbConfig{
			Addr: os.Getenv("LEVELDB"),
		}
	}

	// SSDB 服务配置
	if c.SSDBConfig == nil {
		c.Check("SSDB_ADDR")
		c.Check("SSDB_POOL")
		host, _port, e := net.SplitHostPort(os.Getenv("SSDB_ADDR"))
		port, e := strconv.Atoi(_port)
		if e != nil {
			log.Fatal("SSDB_ADDR 配置异常") // 中断程序时输出
		}
		pools := strings.Split(os.Getenv("SSDB_POOL"), ":")
		if len(pools) != 3 {
			log.Fatal("SSDB_POOL 配置异常") // 中断程序时输出
		}
		minPoolSize, e := strconv.Atoi(pools[0])
		if e != nil {
			log.Fatal("SSDB_POOL 配置异常") // 中断程序时输出
		}
		maxPoolSize, e := strconv.Atoi(pools[1])
		if e != nil {
			log.Fatal("SSDB_POOL 配置异常") // 中断程序时输出
		}
		acquireIncrement, e := strconv.Atoi(pools[2])
		if e != nil {
			log.Fatal("SSDB_POOL 配置异常") // 中断程序时输出
		}
		password := os.Getenv("SSDB_PWD")

		c.SSDBConfig = &conf.Config{
			Host:             host,
			Port:             port,
			MinPoolSize:      minPoolSize,
			MaxPoolSize:      maxPoolSize,
			AcquireIncrement: acquireIncrement,
			Password:         password,
		}
	}

	// nsq 实时消息
	if c.Nsq == nil {
		c.Check("NSQD_ADDR") // 单节点1
		c.Check("NSQC_ADDR")

		c.Nsq = &api_config.NsqConfig{
			NsqdAddr:       os.Getenv("NSQD_ADDR"),
			NsqadminAddr:   os.Getenv("NSQC_HTTP"),
			NsqlookupdAddr: os.Getenv("NSQC_ADDR"),
		}
	}

	log.Printf("加载配置文件并检查配置项: OK\n")
	//log.Printf("%#v \n", JwtConf)
}

// 检查配置项 Must Checked
func (c *ApiConfigs) Check(key string) {
	if _, e := envy.MustGet(key); e != nil {
		panic(e)
	} else {
		//log.Printf("配置文件: %s = %s \n", key, val)
	}
}
