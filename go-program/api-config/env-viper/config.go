package env_viper

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	"github.com/seefan/gossdb/conf"

	"github.com/fsnotify/fsnotify"

	api_config "github.com/angenalZZZ/Go/go-program/api-config"
	"github.com/spf13/viper"
)

var config = &ApiConfigs{} // 当前配置 访问对象

type ApiConfigs struct {
	api_config.ApiConfigs  // 当前配置 所有变量
	api_config.IApiConfigs // 当前配置 需实现的方法
}

// 加载配置文件 插件 github.com/spf13/viper
func init() {
	// 加载配置文件并检查配置项
	config.Load()

	// 当前配置 设置静态变量
	api_config.Config = &config.ApiConfigs
}

// 加载配置文件并检查配置项
func (c *ApiConfigs) Load() {
	p, _ := os.Getwd()
	p = os.Getenv("GOPATH") + "/src/github.com/angenalZZZ/Go/go-program/"

	// 检查环境，判断加载的配置文件 (GO_ENV:环境变量或命令行参数)
	f, v := ".env-prod", os.Getenv("GO_ENV")
	if v == "development" {
		f = ".env-dev"
	} else if v == "test" {
		f = ".env-test"
	}
	//log.Printf("配置文件: 解析 %s%s.yaml\n", p, f)

	viper.SetConfigName(f)      // 指定的配置
	viper.SetConfigType("yaml") // 设置配置文件格式为yaml
	viper.AddConfigPath(p)      // 解析配置文件
	viper.AutomaticEnv()        // 读取匹配的环境变量
	viper.SetEnvPrefix("API")
	//viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	//viper.AllowEmptyEnvVar(true)
	if e := viper.ReadInConfig(); e != nil {
		panic(e)
	} else {
		//log.Printf("配置文件: 解析结果\n %v\n", viper.AllSettings())
	}

	println()

	// 检查配置项目
	if c.Jwt == nil {
		j := "http.auth.jwt."
		c.Check(j + "algorithms")
		c.Check(j + "secret")
		c.Check(j + "lifetime")
		c.Check(j + "issuer")
		c.Check(j + "subject")
		c.Check(j + "audience")
		c.Check(j + "sign.pub")
		c.Check(j + "sign.key")
		c.Check(j + "sign.alg")

		jwtSign := api_config.JwtSign{Key: viper.GetString(j + "sign.key"), Pub: viper.GetString(j + "sign.pub"), Alg: viper.GetString(j + "sign.alg")}

		c.Jwt = &api_config.JwtConfig{
			AUTH_JWT:       "jwt",
			JWT_algorithms: viper.GetString(j + "algorithms"),
			JWT_SECRET:     viper.GetString(j + "secret"),
			JWT_LIFETIME:   viper.GetInt(j + "lifetime"),
			JWT_Issuer:     viper.GetString(j + "issuer"),
			JWT_Subject:    viper.GetString(j + "subject"),
			JWT_Audience:   viper.GetString(j + "audience"),
			JWT_Sign:       jwtSign,
		}
	}

	// HttpWeb服务配置
	if c.HttpSvr == nil {
		j := "http.addr."
		c.Check(j + "host")
		c.Check(j + "post")

		c.HttpSvr = &api_config.HttpSvrConfig{
			Addr: viper.GetString(j+"host") + ":" + viper.GetString(j+"post"),
		}
	}

	// Redis 客户端接口
	if c.RedisCli == nil {
		j := "redis."
		c.Check(j + "addr")
		c.Check(j + "pwd")
		c.Check(j + "db")

		c.RedisCli = &api_config.RedisCliConfig{
			Addr: viper.GetString(j + "addr"),
			Pwd:  viper.GetString(j + "pwd"),
			Db:   viper.GetInt(j + "db"),
		}
		//log.Printf("配置文件: %+v\n", c.RedisCli)
	}

	// LevelDb 存储目录地址
	if c.LevelDb == nil {
		j := "leveldb."
		c.Check(j + "addr")

		c.LevelDb = &api_config.LevelDbConfig{
			Addr: viper.GetString(j + "addr"),
		}
	}

	// SSDB 服务配置
	if c.SSDBConfig == nil {
		j := "ssdb."
		c.Check(j + "addr")
		//c.Check(j + "pwd")
		c.Check(j + "pool.min")
		c.Check(j + "pool.max")
		c.Check(j + "pool.inc")

		host, _port, e := net.SplitHostPort(viper.GetString(j + "addr"))
		port, e := strconv.Atoi(_port)
		if e != nil {
			log.Fatal("SSDB_ADDR 配置异常") // 中断程序时输出
		}

		c.SSDBConfig = &conf.Config{
			Host:             host,
			Port:             port,
			MinPoolSize:      viper.GetInt(j + "pool.min"),
			MaxPoolSize:      viper.GetInt(j + "pool.max"),
			AcquireIncrement: viper.GetInt(j + "pool.inc"),
			Password:         viper.GetString(j + "pwd"),
		}
	}

	// nsq 实时消息
	if c.Nsq == nil {
		c.Check("nsqd") // 单节点1
		c.Check("nsqadmin")
		c.Check("nsqlookupd")

		c.Nsq = &api_config.NsqConfig{
			NsqdAddr:       viper.GetString("nsqd"),
			NsqadminAddr:   viper.GetString("nsqadmin"),
			NsqlookupdAddr: viper.GetString("nsqlookupd"),
		}
	}

	log.Printf("加载配置文件并检查配置项: OK\n")
	//log.Printf("%#v \n", JwtConf)

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("Config file changed: %s", e.Name)
	})

}

// 检查配置项 Must Checked
func (c *ApiConfigs) Check(key string) {
	if viper.IsSet(key) == false {
		panic(fmt.Sprintf("配置文件: %s Not Set In Config", key))
	} else {
		//log.Printf("配置文件: %s = %s \n", key, val)
	}
}
