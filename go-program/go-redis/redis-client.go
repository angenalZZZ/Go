package go_redis

import (
	"angenalZZZ/go-program/api-config"
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

/**
数据库连接 redis : Client
*/
var Db *redis.Client
var op *redis.Options

// 初始化Client
func Init() {
	if Db != nil {
		return
	}

	// config
	api_config.Check("REDIS_ADDR")
	api_config.Check("REDIS_PWD")
	api_config.Check("REDIS_DB")
	i, e := strconv.Atoi(os.Getenv("REDIS_DB"))
	if e != nil {
		i = 0
	}
	op = &redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PWD"),
		DB:       i, // default DB: 0
	}

	// new client
	Db = redis.NewClient(op)

	// check
	if _, e := Db.Ping().Result(); e != nil {
		log.Fatal(e) // 中断程序时输出
	}
}

// 数据库 Redis Client close
func ShutdownClient() {
	log.Println("缓存数据库 Redis Client closing..")
	if e := Db.Close(); e != nil {
		log.Fatal(e) // 中断程序时输出
	}
}

// 测试
func Test() {
	Init()
	log.Printf("缓存数据库 Redis Client testing.. Addr: %s\n\n", op.Addr)

	// redis : Client
	db := Db
	rand.Seed(time.Now().UnixNano())

	// 写入数据 Set
	key, val := fmt.Sprintf("timestamp%d%d", time.Now().Unix(), rand.Intn(1000)), "hello"
	if e := db.Set(key, val, 0).Err(); e != nil {
		log.Printf(" redis Set: Err\n [%s] %v\n", key, e)
	} else {
		log.Printf(" redis Set: Ok\n [%s] %s\n", key, val)
	}

	// 读取数据 Get
	valSaved, e := db.Get(key).Result()
	if e == redis.Nil {
		log.Printf(" redis Get: Nil\n [%s] does not exist\n", key)
	} else if e != nil {
		log.Printf(" redis Get: Err\n [%s] %v\n", key, e)
	} else {
		log.Printf(" redis Get: Ok\n [%s] %s\n", key, valSaved)
	}

	// 写入数据?当key不存在时+过期时间 SET key value EX 10 NX
	key, val = fmt.Sprintf("timestamp%d%d", time.Now().Unix(), rand.Intn(1000)), "values"
	_, e = db.SetNX(key, val, 10*time.Second).Result()
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
		if e = db.LPush(key, v).Err(); e != nil {
			log.Printf(" redis LPush: Err\n [%s] %v\n", key, e)
		} else {
			log.Printf(" redis LPush: Ok\n [%s] %d\n", key, v)
		}
	}
	arr, err := db.Sort(key, &redis.Sort{Offset: 0, Count: 6, Order: "ASC"}).Result()
	if err != nil {
		log.Printf(" redis Sort: Err\n [%s] %v\n", key, e)
	} else {
		log.Printf(" redis Sort: Ok\n [%s] %v\n", key, arr)
	}

	// 读取有序集合中指定分数区间的成员列表 ZRANGEBYSCORE zset0 -inf +inf WITHSCORES LIMIT 0 6 [WITHSCORES:输出分数]
	key = fmt.Sprintf("zset%d%d", time.Now().Unix(), rand.Intn(1000))
	for i := range [6]int{1} {
		v := redis.Z{Score: rand.Float64(), Member: fmt.Sprintf("member%d", rand.Intn(100)+i)}
		// ZADD zset0 1 member1
		if e = db.ZAdd(key, v).Err(); e != nil {
			log.Printf(" redis ZAdd: Err\n [%s] %v\n", key, e)
		} else {
			log.Printf(" redis ZAdd: Ok\n [%s] %v\n", key, v)
		}
	}
	set1, er1 := db.ZRangeByScoreWithScores(key, redis.ZRangeBy{Min: "-inf", Max: "+inf", Offset: 0, Count: 6}).Result()
	if er1 != nil {
		log.Printf(" redis ZRANGEBYSCORE: Err\n [%s] %v\n", key, er1)
	} else {
		log.Printf(" redis ZRANGEBYSCORE: Ok\n [%s] %v\n", key, set1)
	}

	// 计算: 给定有序集的交集,并将该交集(结果集)储存起来 http://www.runoob.com/redis/sorted-sets-zinterstore.html
	// ZINTERSTORE out 2 zset01 zset02 WEIGHTS 2 3 AGGREGATE SUM
	db.ZAddNX("zset01", redis.Z{Score: float64(rand.Intn(100)), Member: "A"}, redis.Z{Score: float64(rand.Intn(100)), Member: "B"})
	db.ZAddNX("zset02", redis.Z{Score: float64(rand.Intn(100)), Member: "A"}, redis.Z{Score: float64(rand.Intn(100)), Member: "B"})
	set2, er2 := db.ZInterStore("zset0102", redis.ZStore{Weights: []float64{0, 100}}, "zset01", "zset02").Result()
	if er1 != nil {
		log.Printf(" redis ZINTERSTORE: Err\n [%s] %v\n", key, er2)
	} else {
		log.Printf(" redis ZINTERSTORE: Ok\n [%s] %v\n", key, set2)
	}

	// 计算: EVAL "return {KEYS[1],ARGV[1]}" 1 "key" "hello"
	//_, er3 := db.Eval("return {KEYS[1],ARGV[1]}", []string{"key"}, "hello").Result()
}
