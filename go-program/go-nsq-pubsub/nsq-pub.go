package go_nsq_pubsub

import (
	"angenalZZZ/go-program/api-config"
	"github.com/nsqio/go-nsq"
	"log"
	"os"
	"time"
)

/**
生产者 nsq Producer : Client
*/
var D *nsq.Producer

// 初始化Cli
func InitProducer() {
	if D != nil {
		return
	}

	// config
	api_config.Check("NSQD_ADDR") // 单节点1
	addr := os.Getenv("NSQD_ADDR")
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
