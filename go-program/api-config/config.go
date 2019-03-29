package api_config

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/seefan/gossdb/conf"
)

// 当前配置访问对象
var (
	Config *ApiConfigs
	Check  func(key string) // 检查配置项 Must Checked
)

// 当前配置访问对象 需实现的方法
type IApiConfigs interface {
	LoadCheck()
	Check(key string)
}

type ApiConfigs struct {
	// HttpWeb 服务配置
	HttpSvr *HttpSvrConfig

	// 认证方式 Jwt Config
	Jwt *JwtConfig

	// Redis 服务配置
	RedisCli *RedisCliConfig

	// LevelDb 存储目录地址
	LevelDb *LevelDbConfig

	// SSDB 服务配置
	SSDBConfig *conf.Config

	// nsq 实时消息
	Nsq *NsqConfig

	IApiConfigs
}

// HttpWeb 服务配置
type HttpSvrConfig struct {
	Addr string // TCP address to listen on, ":http" if empty
}

// Redis 服务配置
type RedisCliConfig struct {
	Addr string // TCP address
	Pwd  string
	Db   int
}

// LevelDb 存储目录地址
type LevelDbConfig struct {
	Addr string
}

// nsq 实时消息
type NsqConfig struct {
	NsqdAddr string

	NsqadminAddr string

	NsqlookupdAddr string
}

// 认证方式 Jwt Config: 复合类型
type JwtConfig struct {
	AUTH_JWT, JWT_algorithms, JWT_SECRET  string
	JWT_LIFETIME                          int
	JWT_Issuer, JWT_Subject, JWT_Audience string
	JWT_Sign                              JwtSign
}
type JwtSign struct {
	Key, Pub, Alg string
}

func (o *JwtSign) HasKeyAndPub() bool {
	return o.Key != "" && o.Pub != "" && o.Alg != ""
}

// 加载输入参数|文件路径
func LoadArgInput(s string) ([]byte, error) {
	if s == "" {
		return nil, fmt.Errorf("no input")
	} else if s == "+" {
		return []byte("{}"), nil
	}
	var r io.Reader
	if s == "-" {
		r = os.Stdin
	} else {
		if f, e := os.Open(s); e != nil {
			return nil, e
		} else {
			defer f.Close() // end
			r = f
		}
	}
	return ioutil.ReadAll(r)
}

// "+" > {}
func JsonParse(s string) (v interface{}, e error) {
	var data []byte
	if data, e = LoadArgInput(s); e == nil {
		e = json.Unmarshal(data, v)
	}
	return
}

// {} > "+"
func JsonStringify(v interface{}, indent bool) (s []byte, e error) {
	if indent == false {
		s, e = json.MarshalIndent(v, "", "    ")
	} else {
		s, e = json.Marshal(v)
	}
	return
}
