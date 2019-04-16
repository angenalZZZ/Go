package go_redis

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/gomodule/redigo/redis"
)

/************ List（列表）*************/
func TestCli_list(c redis.Conn) {

	// 写入数据 Set
	key := fmt.Sprintf("list%d%d", time.Now().Unix(), rand.Intn(1000))
	for i := range [6]int{1} {
		v := rand.Intn(1000) + i
		// LPUSH list0 1
		if _, e := c.Do("LPUSH", key, v); e != nil {
			log.Printf(" redis LPush: Err\n [%s] %v\n", key, e)
		} else {
			log.Printf(" redis LPush: Ok\n [%s] %d\n", key, v)
		}
	}

	// 读取数据 Get 前 2 条
	if valSaved, e := c.Do("LRANGE", key, 0, 2); valSaved == nil {
		log.Printf(" redis Get: Nil\n [%s] does not exist\n", key)
	} else if e != nil {
		log.Printf(" redis Get: Err\n [%s] %v\n", key, e)
	} else {
		log.Printf(" redis Get: Ok\n [%s] %s\n", key, valSaved)
	}

	// 读取数组+分页+排序 SORT list0 LIMIT 0 6 ASC, sort list0 desc alpha
	if arr, e := c.Do("SORT", key, "LIMIT", 0, 6, "ASC"); e != nil {
		log.Printf(" redis Sort: Err\n [%s] %v\n", key, e)
	} else {
		log.Printf(" redis Sort: Ok\n [%s] %s\n", key, arr)
	}

	// 删除数据 Del
	if _, e := c.Do("DEL", key); e != nil {
		log.Printf(" redis Del: Err\n [%s] %v\n", key, e)
	} else {
		log.Printf(" redis Del: Ok\n [%s]\n", key)
	}

}
