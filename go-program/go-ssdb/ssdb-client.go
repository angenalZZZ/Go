package go_ssdb

import (
	"fmt"
	"github.com/angenalZZZ/Go/go-program/api-config"
	"github.com/seefan/gossdb"
	"github.com/seefan/gossdb/conf"
	"github.com/seefan/gossdb/ssdb"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

/**
数据库连接 SSdb : Client
*/
var Db *gossdb.Client
var op *conf.Config
var pool *gossdb.Connectors

// 初始化配置
func init() {
	// config
	api_config.Check("SSDB_ADDR")
	api_config.Check("SSDB_POOL")
	host, _port, e := net.SplitHostPort(os.Getenv("SSDB_ADDR"))
	port, e := strconv.Atoi(_port)
	if e != nil {
		log.Fatal("SSDB_ADDR 配置异常") // 中断程序时输出
	}
	pools := strings.Split(os.Getenv("SSDB_POOL"), ":")
	if len(pools) != 3 {
		log.Fatal("SSDB_POOL 配置异常") // 中断程序时输出
	}
	minPoolSize, e := strconv.Atoi(pools[0])
	if e != nil {
		log.Fatal("SSDB_POOL 配置异常") // 中断程序时输出
	}
	maxPoolSize, e := strconv.Atoi(pools[1])
	if e != nil {
		log.Fatal("SSDB_POOL 配置异常") // 中断程序时输出
	}
	acquireIncrement, e := strconv.Atoi(pools[2])
	if e != nil {
		log.Fatal("SSDB_POOL 配置异常") // 中断程序时输出
	}
	password := os.Getenv("SSDB_PWD")

	op = &conf.Config{
		Host:             host,
		Port:             port,
		MinPoolSize:      minPoolSize,
		MaxPoolSize:      maxPoolSize,
		AcquireIncrement: acquireIncrement,
		Password:         password,
	}
}
func initDb() {
	if Db != nil {
		return
	}

	// new Pool
	var e error
	pool, e = gossdb.NewPool(op)
	if e != nil {
		log.Fatal(e) // 中断程序时输出
	}
	if e = pool.Start(); e != nil {
		log.Fatal(e) // 中断程序时输出
	}

	// new client
	Db, e = pool.NewClient()
	if e != nil {
		log.Fatal(e) // 中断程序时输出
	}

	// check
	if e := Db.Ping(); e == false {
		log.Fatal("SSDB_ADDR 配置异常-Ping-失败") // 中断程序时输出
	}
}

// new client from pool, defer db.Close()
func NewClient() *gossdb.Client {
	db, _ := pool.NewClient()
	return db
}

// 数据库 SSdb Client close
func ShutdownClient() {
	if Db != nil && pool != nil {
		//log.Println("缓存数据库 SSdb Client closing..")
		if e := Db.Close(); e != nil {
			log.Fatal(e) // 中断程序时输出
		}
		pool.Close()
		ssdb.Close()
	}
}

// 测试
func Test() {
	initDb()
	log.Printf("缓存数据库 SSdb Client testing.. Addr: %s:%d\n\n", op.Host, op.Port)

	// SSdb : Client
	db := Db
	rand.Seed(time.Now().UnixNano())
	ttl := int64(60) // 设置过期(秒)

	// 写入数据 Set
	key, val := fmt.Sprintf("timestamp%d%d", time.Now().Unix(), rand.Intn(1000)), "hello"
	if e := db.Set(key, val, ttl); e != nil {
		log.Printf(" SSdb Set: Err\n [%s] %v\n", key, e)
	} else {
		log.Printf(" SSdb Set: Ok\n [%s] %s\n", key, val)
	}

	// 读取数据 Get
	valSaved, e := db.Get(key)
	if e != nil {
		log.Printf(" SSdb Get: Err\n [%s] %v\n", key, e)
	} else {
		log.Printf(" SSdb Get: Ok\n [%s] %s\n", key, valSaved)
	}

	// 简单使用
	func() {
		// new pool
		if err := ssdb.Start(op); err != nil {
			log.Printf(" 简单使用 SSdb Start: 无法连接到SSdb")
		} else {
			log.Printf(" 简单使用 SSdb Start..  Addr: %s:%d\n", op.Host, op.Port)
		}
		defer ssdb.Close()

		// new client
		c, e := ssdb.Client()
		if e != nil {
			log.Printf(" 简单使用 SSdb 无法获取连接from-SSdb-pool")
		} else {
			log.Printf(" 简单使用 SSdb 获取连接from-SSdb-pool: OK\n")
		}
		defer c.Close()

		//some simple run
		c.Set("a", 1)
		c.Get("a")
		c.Del("a")
		log.Printf(" 简单使用 SSdb Set/Get/Del: OK\n")

		//another simple run
		ssdb.Simple(func(c *gossdb.Client) (e error) {
			e = c.Set("test", "hello world")
			if e == nil {
				if _, e = c.Get("test"); e == nil {
					if e = c.Del("test"); e == nil {
						log.Printf(" 简单使用 SSdb /Simple/ Set/Get/Del: OK\n")
					}
				}
			}
			if e != nil {
				log.Printf(" 简单使用 SSdb /Simple/ error")
			}
			return
		})

	}()

}
