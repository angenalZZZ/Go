package go_nsq_pubsub

import (
	"log"
	"time"

	api_config "github.com/angenalZZZ/Go/go-program/api-config"
	"github.com/nsqio/go-nsq"
)

/**
生产者 nsq Producer : Client
*/
var D *nsq.Producer

// 初始化配置
func InitProducer() {
	// config
	// 单节点1
	addr := api_config.Config.Nsq.NsqdAddr
	config := nsq.NewConfig()
	config.WriteTimeout = 3 * time.Second

	// client
	d, e := nsq.NewProducer(addr, config)
	if e != nil {
		log.Fatal(e) // 中断程序时输出
	}
	D = d

	// check
	if e := D.Ping(); e != nil {
		log.Fatal(e) // 中断程序时输出
	}
}
