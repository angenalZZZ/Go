package go_nsq_pubsub

import (
	"angenalZZZ/go-program/api-config"
	"github.com/nsqio/go-nsq"
	"log"
	"os"
	"time"
)

/**
消费者 nsq Producer : Client
*/
var C *nsq.Consumer

// 初始化Cli
func InitConsumer() {
	if C != nil {
		return
	}

	// config
	api_config.Check("NSQC_ADDR") // 单节点
	addr := os.Getenv("NSQC_ADDR")
	config := nsq.NewConfig()
	config.ReadTimeout = 3 * time.Second
	config.LookupdPollInterval = 2 * time.Second // 设置心跳

	// client
	c, e := nsq.NewConsumer("client1", "channel1", config)
	if e != nil {
		log.Fatal(e) // 中断程序时输出
	}
	c.AddHandler(nsq.HandlerFunc(handle))
	C = c

	// check
	if e := C.ConnectToNSQLookupd(addr); e != nil {
		log.Fatal(e) // 中断程序时输出
	}
}

// handle message
func handle(m *nsq.Message) error {
	// handle the message
	log.Printf("receive ID:%s,addr:%s,message:%s", m.ID, m.NSQDAddress, m.Body)
	return nil
}