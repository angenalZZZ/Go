package go_redis

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	go_type "github.com/angenalZZZ/Go/go-program/go-type"

	api_config "github.com/angenalZZZ/Go/go-program/api-config"

	"github.com/gomodule/redigo/redis"
)

/**
数据库连接 redis : Client
*/
var Cli redis.Conn
var CliPoll *redis.Pool
var cliOpt redis.DialOption
var cliAddr string

// 初始化配置
func init() {
	// config
	cliAddr = api_config.Config.RedisCli.Addr
	// client
	cliOpt = redis.DialClientName("redis-cli")
	cliOpt = redis.DialUseTLS(false)
	// password
	password := api_config.Config.RedisCli.Pwd
	if len(password) > 0 {
		cliOpt = redis.DialPassword(password)
	}
	// db number
	cliOpt = redis.DialDatabase(api_config.Config.RedisCli.Db)
}
func initCli() {
	if CliPoll != nil {
		return
	}

	// managed Pool
	CliPoll = &redis.Pool{
		MaxActive: 5,
		MaxIdle:   5,
		Wait:      true,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", cliAddr, cliOpt)
		},
	}

	// new client
	Cli = CliPoll.Get()

	// check
	if e := Cli.Err(); e != nil {
		log.Fatal(e) // 中断程序时输出
	}
}

// 数据库 Redis Cli close
func ShutdownCli() {
	if CliPoll != nil {
		//log.Println("缓存数据库 Redis Cli closing..")
		if Cli != nil {
			if e := Cli.Close(); e != nil {
				log.Fatal(e) // 中断程序时输出
			}
		}
		if e := CliPoll.Close(); e != nil {
			log.Fatal(e) // 中断程序时输出
		}
	}
}

// 测试
func TestCli() {
	initCli()
	log.Printf("缓存数据库 Redis Cli testing.. Addr: %s\n\n", cliAddr)

	// redis : new Cli
	c := CliPoll.Get()
	defer func() { _ = c.Close() }()
	rand.Seed(time.Now().UnixNano())

	/************ String（字符串）*************/

	// 写入数据 Set
	key, val := fmt.Sprintf("timestamp%d%d", time.Now().Unix(), rand.Intn(1000)), "hello"
	if _, e := c.Do("SET", key, val); e != nil {
		log.Printf(" redis Set: Err\n [%s] %v\n", key, e)
	} else {
		log.Printf(" redis Set: Ok\n [%s] %s\n", key, val)
	}

	// 读取数据 Get
	valSaved, e := c.Do("GET", key)
	if valSaved == nil {
		log.Printf(" redis Get: Nil\n [%s] does not exist\n", key)
	} else if e != nil {
		log.Printf(" redis Get: Err\n [%s] %v\n", key, e)
	} else {
		log.Printf(" redis Get: Ok\n [%s] %s\n", key, valSaved)
	}

	// 删除数据 Del
	_, e = c.Do("DEL", key)
	if e != nil {
		log.Printf(" redis Del: Err\n [%s] %v\n", key, e)
	} else {
		log.Printf(" redis Del: Ok\n [%s]\n", key)
	}

	/************ Hash（哈希）*************/
	keyHash, valHash := fmt.Sprintf("hash%d%d", time.Now().Unix(), rand.Intn(1000)), go_type.Q{
		"a": 1,
		"b": "2b",
	}.Slice()
	_, e = c.Do("HMSET", keyHash, valHash...)
	if e != nil {
		log.Printf(" redis HMSET: Err\n [%s]\n", keyHash)
	} else {
		log.Printf(" redis HMSET: Ok\n [%s]\n", keyHash)
	}
	_, e = c.Do("DEL", keyHash)

	// 写入数据?当key不存在时+过期时间 SET key value EX 10 NX
	key, val = fmt.Sprintf("timestamp%d%d", time.Now().Unix(), rand.Intn(1000)), "values"
	_, e = c.Do("SET", key, val, "EX", 10, "NX")
	if e != nil {
		log.Printf(" redis SetNX: Err\n [%s] %v\n", key, e)
	} else {
		log.Printf(" redis SetNX: Ok\n [%s] %s\n", key, val)
	}

	// 读取数组+分页+排序 SORT list0 LIMIT 0 6 ASC, sort list0 desc alpha
	key = fmt.Sprintf("list%d%d", time.Now().Unix(), rand.Intn(1000))
	for i := range [6]int{1} {
		v := rand.Intn(1000) + i
		// LPUSH list0 1
		if _, e = c.Do("LPUSH", key, v); e != nil {
			log.Printf(" redis LPush: Err\n [%s] %v\n", key, e)
		} else {
			log.Printf(" redis LPush: Ok\n [%s] %d\n", key, v)
		}
	}
	arr, err := c.Do("SORT", key, "LIMIT", 0, 6, "ASC")
	if err != nil {
		log.Printf(" redis Sort: Err\n [%s] %v\n", key, e)
	} else {
		log.Printf(" redis Sort: Ok\n [%s] %v\n", key, arr)
	}

	// 读取有序集合中指定分数区间的成员列表 ZRANGEBYSCORE zset0 -inf +inf WITHSCORES LIMIT 0 6 [WITHSCORES:输出分数]
	key = fmt.Sprintf("zset%d%d", time.Now().Unix(), rand.Intn(1000))
	for i := range [6]int{1} {
		score, member := rand.Float64(), fmt.Sprintf("member%d", rand.Intn(100)+i)
		// ZADD zset0 1 member1
		if _, e = c.Do("ZADD", key, score, member); e != nil {
			log.Printf(" redis ZAdd: Err\n [%s] %v\n", key, e)
		} else {
			log.Printf(" redis ZAdd: Ok\n [%s] %s=%f\n", key, member, score)
		}
	}
	set1, er1 := c.Do("ZRANGEBYSCORE", key, "-inf", "+inf", "WITHSCORES", "LIMIT", 0, 6)
	if er1 != nil {
		log.Printf(" redis ZRANGEBYSCORE: Err\n [%s] %v\n", key, er1)
	} else {
		log.Printf(" redis ZRANGEBYSCORE: Ok\n [%s] %v\n", key, set1)
	}

	// 计算: 给定有序集的交集,并将该交集(结果集)储存起来 http://www.runoob.com/redis/sorted-sets-zinterstore.html
	// ZINTERSTORE out 2 zset01 zset02 WEIGHTS 2 3 AGGREGATE SUM
	e = c.Send("ZADD", "zset01", rand.Intn(100), "A", "NX")
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
	// 交集的目标key
	key = "zset0102"
	set2, er2 := c.Do("ZINTERSTORE", key, 2, "zset01", "zset02", "WEIGHTS", 0, 100, "AGGREGATE", "SUM")
	if er1 != nil {
		log.Printf(" redis ZINTERSTORE: Err\n [%s] %v\n", key, er2)
	} else {
		log.Printf(" redis ZINTERSTORE: Ok\n [%s] %v\n", key, set2)
	}

	// 计算: EVAL "return {KEYS[1],ARGV[1]}" 1 "key" "hello"
	//_, er3 := c.Do("EVAL", "return {KEYS[1],ARGV[1]}", 1, "key", "hello")
}
