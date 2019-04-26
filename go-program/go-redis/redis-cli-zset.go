package go_redis

import (
	"fmt"
	"math/rand"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/gomodule/redigo/redis"
)

/************ ZSet（有序集合）*************/
func DoCli_zset(c redis.Conn) {

	// 写入数据 Set

	// 读取有序集合中指定分数区间的成员列表 ZRANGEBYSCORE zset0 -inf +inf WITHSCORES LIMIT 0 6 [WITHSCORES:输出分数]
	key := fmt.Sprintf("zset%d%d", time.Now().Unix(), rand.Intn(1000))
	for i := range [6]int{1} {
		score, member := rand.Float64(), fmt.Sprintf("member%d", rand.Intn(100)+i)
		// ZADD zset0 1 member1
		if _, e := c.Do("ZADD", key, score, member); e != nil {
			log.Printf(" redis ZAdd: Err\n [%s] %v\n", key, e)
		} else {
			log.Printf(" redis ZAdd: Ok\n [%s] %s=%f\n", key, member, score)
		}
	}

	// 读取数据 Get
	if set1, er1 := c.Do("ZRANGEBYSCORE", key, "-inf", "+inf", "WITHSCORES", "LIMIT", 0, 6); er1 != nil {
		log.Printf(" redis ZRANGEBYSCORE: Err\n [%s] %v\n", key, er1)
	} else {
		log.Printf(" redis ZRANGEBYSCORE: Ok\n [%s] %s\n", key, set1)
	}

	// 删除数据 Del
	if _, e := c.Do("DEL", key); e != nil {
		log.Printf(" redis Del: Err\n [%s] %v\n", key, e)
	} else {
		log.Printf(" redis Del: Ok\n [%s]\n", key)
	}

	// 计算: 给定有序集的交集,并将该交集(结果集)储存起来 http://www.runoob.com/redis/sorted-sets-zinterstore.html
	// ZINTERSTORE out 2 zset01 zset02 WEIGHTS 2 3 AGGREGATE SUM
	e := c.Send("ZADD", "zset01", rand.Intn(100), "A", "NX")
	e = c.Send("ZADD", "zset01", rand.Intn(100), "B", "NX")
	e = c.Send("ZADD", "zset01", rand.Intn(100), "C", "NX")
	e = c.Send("ZADD", "zset02", rand.Intn(100), "A", "NX")
	e = c.Send("ZADD", "zset02", rand.Intn(100), "B", "NX")
	e = c.Send("ZADD", "zset02", rand.Intn(100), "C", "NX")
	e = c.Flush()
	_, e = c.Receive()
	_, e = c.Receive()
	_, e = c.Receive()
	_, e = c.Receive()
	_, e = c.Receive()
	_, e = c.Receive()

	if e != nil {
		key = "zset0102" // 交集的目标key
		if set2, er2 := c.Do("ZINTERSTORE", key, 2, "zset01", "zset02", "WEIGHTS", 0, 100, "AGGREGATE", "SUM"); er2 != nil {
			log.Printf(" redis ZINTERSTORE: Err\n [%s] %v\n", key, er2)
		} else {
			log.Printf(" redis ZINTERSTORE: Ok\n [%s] %s\n", key, set2)
		}
	}

}
