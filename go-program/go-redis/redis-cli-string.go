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

	// 时间戳：以秒计
	timestamp := time.Now().Unix()
	// 时间戳：以毫秒计
	timestampNano := time.Now().UnixNano()

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

	// 设置过期时间，以秒计
	c.Do("expire", key, 60)
	c.Do("expireat", key, 60+timestamp)
	// 设置过期时间，以毫秒计
	c.Do("pexpire", key, 60000)
	c.Do("pexpireat", key, 60000+timestampNano)
	// 获取过期时间，以秒为单位
	if t, e := c.Do("ttl", key); e != nil {
		log.Printf(" redis ttl: Err\n [%s] %v\n", key, e)
	} else {
		log.Printf(" redis ttl: Ok\n [%s] = %d\n", key, t)
	}
	// 获取过期时间，以毫秒为单位
	if t, e := c.Do("pttl", key); e != nil {
		log.Printf(" redis pttl: Err\n [%s] %v\n", key, e)
	} else {
		log.Printf(" redis pttl: Ok\n [%s] = %d\n", key, t)
	}
	// 移除 key 的过期时间，key 将持久保持
	if _, e := c.Do("persist", key); e != nil {
		log.Printf(" redis persist: Err\n [%s] %v\n", key, e)
	} else {
		log.Printf(" redis persist: Ok\n [%s]\n", key)
	}

	// 删除数据 Del
	if _, e := c.Do("DEL", key); e != nil {
		log.Printf(" redis Del: Err\n [%s] %v\n", key, e)
	} else {
		log.Printf(" redis Del: Ok\n [%s]\n", key)
	}

	// 检查给定 key 是否存在
	if nx, e := c.Do("exists", key); e != nil {
		log.Printf(" redis exists: Err\n [%s] %v\n", key, e)
	} else {
		log.Printf(" redis exists: Ok\n [%s] = %v\n", key, nx)
	}

}
