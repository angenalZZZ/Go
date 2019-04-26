package go_redis

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/gomodule/redigo/redis"
)

/************ Set（集合）*************/
func DoCli_set(c redis.Conn) {

	// 写入数据 Set
	key, val := fmt.Sprintf("hash%d%d", time.Now().Unix(), rand.Intn(1000)), []string{"a", "b"}
	for _, k := range val {
		if _, e := c.Do("sadd", key, k); e != nil {
			log.Printf(" redis sadd: Err\n [%s][%s]\n", key, k)
		} else {
			log.Printf(" redis sadd: Ok\n [%s][%s]\n", key, k)
		}
	}

	// 读取数据 Get
	if valSaved, e := c.Do("smembers", key); valSaved == nil {
		log.Printf(" redis smembers: Nil\n [%s] does not exist\n", key)
	} else if e != nil {
		log.Printf(" redis smembers: Err\n [%s] %v\n", key, e)
	} else {
		log.Printf(" redis smembers: Ok\n [%s] %v\n", key, valSaved)
	}

	// 删除数据 Del
	if _, e := c.Do("DEL", key); e != nil {
		log.Printf(" redis Del: Err\n [%s] %v\n", key, e)
	} else {
		log.Printf(" redis Del: Ok\n [%s]\n", key)
	}

}
