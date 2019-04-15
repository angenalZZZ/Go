package go_redis

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/gomodule/redigo/redis"
)

/************ String（字符串）*************/
func TestCli_string(c redis.Conn) {

	// 写入数据 Set
	key, val := fmt.Sprintf("timestamp%d%d", time.Now().Unix(), rand.Intn(1000)), "hello"
	if _, e := c.Do("SET", key, val); e != nil {
		log.Printf(" redis Set: Err\n [%s] %v\n", key, e)
	} else {
		log.Printf(" redis Set: Ok\n [%s] %s\n", key, val)
	}

	// 写入数据?当key不存在时+过期时间 SET key value EX 60 NX (缓存1分钟)
	key, val = fmt.Sprintf("timestamp%d%d", time.Now().Unix(), rand.Intn(1000)), "values"
	if _, e := c.Do("SET", key, val, "EX", 60, "NX"); e != nil {
		log.Printf(" redis SetNX: Err\n [%s] %v\n", key, e)
	} else {
		log.Printf(" redis SetNX: Ok\n [%s] %s\n", key, val)
	}

	// 读取数据 Get
	if valSaved, e := c.Do("GET", key); valSaved == nil {
		log.Printf(" redis Get: Nil\n [%s] does not exist\n", key)
	} else if e != nil {
		log.Printf(" redis Get: Err\n [%s] %v\n", key, e)
	} else {
		log.Printf(" redis Get: Ok\n [%s] %s\n", key, valSaved)
	}

	// 删除数据 Del
	if _, e := c.Do("DEL", key); e != nil {
		log.Printf(" redis Del: Err\n [%s] %v\n", key, e)
	} else {
		log.Printf(" redis Del: Ok\n [%s]\n", key)
	}

}
