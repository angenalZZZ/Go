package api_config

import (
	"github.com/seefan/gossdb/conf"
)

// 当前配置 设置静态变量
var Config *ApiConfigs

// 当前配置 需实现的方法
type IApiConfigs interface {
	Load()
	Check(key string)
}

// 当前配置 所有变量
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
