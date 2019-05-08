package go_redis

import (
	"fmt"
	"math/rand"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/angenalZZZ/Go/go-program/pkg/types"

	"github.com/gomodule/redigo/redis"
)

/************ Hash（哈希）*************/
func DoCli_hash(c redis.Conn) {

	// 写入数据 Set
	key, val := fmt.Sprintf("hash%d%d", time.Now().Unix(), rand.Intn(1000)), types.Q{
		"a": 1,
		"b": "2",
	}.Slice()
	kv := append([]interface{}{key}, val...)
	if _, e := c.Do("HMSET", kv...); e != nil {
		log.Printf(" redis HMSET: Err\n [%s]\n", key)
	} else {
		log.Printf(" redis HMSET: Ok\n [%s] = %v\n", key, kv[1:])
	}

	// 读取数据 Get
	if valSaved, e := c.Do("HGET", key, val[0]); valSaved == nil {
		log.Printf(" redis HGET: Nil\n [%s][%s] does not exist\n", key, val[0])
	} else if e != nil {
		log.Printf(" redis HGET: Err\n [%s][%s] %v\n", key, val[0], e)
	} else {
		log.Printf(" redis HGET: Ok\n [%s][%s] %s\n", key, val[0], valSaved)
	}

	// 删除数据 Del
	if _, e := c.Do("DEL", key); e != nil {
		log.Printf(" redis Del: Err\n [%s] %v\n", key, e)
	} else {
		log.Printf(" redis Del: Ok\n [%s]\n", key)
	}

}
