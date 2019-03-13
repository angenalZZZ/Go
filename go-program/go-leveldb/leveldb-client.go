package go_leveldb

import (
	"angenalZZZ/go-program/api-config"
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"log"
	"math/rand"
	"os"
	"time"
)

/**
数据库连接 Leveldb : Client
*/
var Db *leveldb.DB
var op *opt.Options
var addr string

// 初始化Client
func Init() {
	if Db != nil {
		return
	}

	// config
	api_config.Check("LEVELDB")
	addr = os.Getenv("LEVELDB")
	//op = &opt.Options{}

	db, e := leveldb.OpenFile(addr, op)
	if e != nil {
		log.Fatal(e) // 中断程序时输出
	}
	Db = db
}

// 数据库 Leveldb Client close
func ShutdownClient() {
	log.Println("内存数据库 Leveldb closing..")
	if e := Db.Close(); e != nil {
		log.Fatal(e) // 中断程序时输出
	}
}

// 扩展 Leveldb 方法
func Put(key, value string, wo *opt.WriteOptions) error {
	return Db.Put([]byte(key), []byte(value), wo)
}
func Get(key string, ro *opt.ReadOptions) (value string, err error) {
	if v, e := Db.Get([]byte(key), ro); e == nil {
		value, err = string(v), e
	} else {
		value, err = "", e
	}
	return
}

// 测试
func Test() {
	Init()
	log.Printf("内存数据库 Leveldb Client testing.. Addr: %s\n\n", addr)

	rand.Seed(time.Now().UnixNano())

	// 写入数据 Put
	key, val := fmt.Sprintf("timestamp%d%d", time.Now().Unix(), rand.Intn(1000)), "hello"
	if e := Put("", "", nil); e != nil {
		log.Printf(" leveldb Put: Err\n [%s] %v\n", key, e)
	} else {
		log.Printf(" leveldb Put: Ok\n [%s] %s\n", key, val)
	}

	// 读取数据 Get
	valSaved, e := Get(key, nil)
	if e != nil {
		log.Printf(" leveldb Get: Err\n [%s] %v\n", key, e)
	} else {
		log.Printf(" leveldb Get: Ok\n [%s] %s\n", key, valSaved)
	}

}
