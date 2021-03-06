package go_redis

import (
	"fmt"
	"math/rand"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/gomodule/redigo/redis"
)

/************ String（字符串）*************/
func DoCli_string(c redis.Conn) {

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

	// 写入数据?当key不存在时+过期时间 SET key value EX 60 NX (缓存1分钟) = SETEX + SETNX
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
	newValue := "new value!"
	if oldSaved, e := c.Do("GETSET", key, newValue); oldSaved == nil {
		log.Printf(" redis GetSet: Nil\n [%s] does not exist\n", key)
	} else if e != nil {
		log.Printf(" redis GetSet: Err\n [%s] %v\n", key, e)
	} else {
		log.Printf(" redis GetSet: Ok\n [%s] %s > %s\n", key, oldSaved, newValue)
	}
	// 读取数据 子字符 GETRANGE key start end
	// 1.读取首字母｜尾字母
	func(key string) {
		// 开始一个事务
		_ = c.Send("MULTI")
		_ = c.Send("GETRANGE", key, 0, 0)
		_ = c.Send("GETRANGE", key, -1, -1)
		// 触发并执行事务
		if valSaved, e := redis.Values(c.Do("EXEC")); valSaved == nil {
			log.Printf(" redis Gets: Nil\n [%s] does not exist\n", key)
		} else if e != nil {
			log.Printf(" redis Gets: Err\n [%s] %v\n", key, e)
		} else {
			log.Printf(" redis Gets: Ok\n [%s] %s\n", key, valSaved)
		}
	}(key)
	// 2.获取指定偏移量上的位(bit)
	func(key string, offset int) {
		valSaved, _ := c.Do("GETBIT", key, offset)
		log.Printf(" redis GetBit: Ok\n [%s][%d], %v %[3]T\n", key, offset, valSaved)
		if v, OK := valSaved.(int64); OK {
			v = 1 - v // 修改位数据offset
			_, _ = c.Do("SETBIT", key, offset, v)
			valSaved, _ = c.Do("GET", key)
			log.Printf(" redis GetBit: Ok, after SetBit\n [%s] %s\n", key, valSaved)
		}
	}(key, 28)

	// 设置过期时间，以秒计
	_, _ = c.Do("expire", key, 60)
	_, _ = c.Do("expireat", key, 60+timestamp)
	_, _ = c.Do("setex", key, 60, "new value")
	// 设置过期时间，以毫秒计
	_, _ = c.Do("pexpire", key, 60000)
	_, _ = c.Do("pexpireat", key, 60000+timestampNano)
	// 获取过期时间，以秒为单位 ttl = -1 表示无过期时间，持久保持
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
	// 移除 key 的过期时间，key 将持久保持 ttl = -1
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
