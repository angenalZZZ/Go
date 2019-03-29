package go_nsq_pubsub

import (
	"log"
	"time"

	api_config "github.com/angenalZZZ/Go/go-program/api-config"
	"github.com/nsqio/go-nsq"
)

/**
消费者 nsq Consumer : Client
*/
var C *nsq.Consumer

// 初始化配置
func init() {
	// config
	addr := api_config.Config.Nsq.NsqlookupdAddr
	config := nsq.NewConfig()
	config.ReadTimeout = 3 * time.Second
	config.LookupdPollInterval = 2 * time.Second // 设置心跳

	// client
	c, e := nsq.NewConsumer("client1", "channel1", config)
	if e != nil {
		log.Fatal(e) // 中断程序时输出
	}
	c.AddHandler(nsq.HandlerFunc(doPrint))
	C = c

	// check
	if e := C.ConnectToNSQLookupd(addr); e != nil {
		log.Fatal(e) // 中断程序时输出
	}
}

// handle message
func doPrint(m *nsq.Message) error {
	// handle the message
	log.Printf("receive ID:%s,addr:%s,message:%s", m.ID, m.NSQDAddress, m.Body)
	return nil
}
