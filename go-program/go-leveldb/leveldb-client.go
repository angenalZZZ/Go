package go_leveldb

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	api_config "github.com/angenalZZZ/Go/go-program/api-config"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"github.com/syndtr/goleveldb/leveldb/util"
)

/**
数据库连接 Leveldb : Client
*/
var Db *leveldb.DB
var op *opt.Options
var addr string

// 初始化配置
func init() {
	// config
	addr = api_config.Config.LevelDb.Addr
	//op = &opt.Options{}
}
func initDb() {
	if Db != nil {
		return
	}

	db, e := leveldb.OpenFile(addr, op)
	if e != nil {
		log.Fatal(e) // 中断程序时输出
	}
	Db = db
}

// 数据库 Leveldb Client close
func ShutdownClient() {
	if Db != nil {
		//log.Println("内存数据库 Leveldb Client closing..")
		if e := Db.Close(); e != nil {
			log.Fatal(e) // 中断程序时输出
		}
	}
}

///数据库方法////////////////////////////////////////////////////////
// 写入数据
func Put(key, value string, wo *opt.WriteOptions) error {
	return Db.Put([]byte(key), []byte(value), wo)
}

// 读取数据
func Get(key string, ro *opt.ReadOptions) (value string, err error) {
	if v, e := Db.Get([]byte(key), ro); e == nil {
		value, err = string(v), e
	} else {
		value, err = "", e
	}
	return
}

// 删除数据
func Del(key string, wo *opt.WriteOptions) error {
	return Db.Delete([]byte(key), wo)
}

// 遍历全部 keys
func RangeAll(ro *opt.ReadOptions, iterator func(key, value []byte, err error)) {
	iter := Db.NewIterator(nil, ro)
	defer iter.Release()
	for iter.Next() {
		iterator(iter.Key(), iter.Value(), iter.Error())
	}
}

// 遍历范围 keys: start ~ limit
func RangeWith(start, limit string, ro *opt.ReadOptions, iterator func(key, value []byte, err error)) {
	iter := Db.NewIterator(&util.Range{Start: []byte(start), Limit: []byte(limit)}, ro)
	defer iter.Release()
	for iter.Next() {
		iterator(iter.Key(), iter.Value(), iter.Error())
	}
}

// 遍历范围 keys: Start With prefix
func RangeStartWith(prefix string, ro *opt.ReadOptions, iterator func(key, value []byte, err error)) {
	iter := Db.NewIterator(util.BytesPrefix([]byte(prefix)), ro)
	defer iter.Release()
	for iter.Next() {
		iterator(iter.Key(), iter.Value(), iter.Error())
	}
}

///测试//////////////////////////////////////////////////////////////
func Do() {
	initDb()
	log.Printf("内存数据库 Leveldb Client testing.. Addr: %s\n\n", addr)

	rand.Seed(time.Now().UnixNano())
	timestamp := time.Now().Unix()

	// 写入数据 Put
	key, val := fmt.Sprintf("timestamp%d%d", timestamp, rand.Intn(1000)), "hello"
	if e := Put(key, val, nil); e != nil {
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

	// 删除数据 Del
	if e = Del(key, nil); e != nil {
		log.Printf(" leveldb Del: Err\n [%s] %v\n", key, e)
	} else {
		log.Printf(" leveldb Del: Ok\n [%s]\n", key)
	}

	// 遍历 keys,values
	for i := range [6]int{1} {
		v := rand.Intn(1000) + i
		// 写入数据 Put
		key, val := fmt.Sprintf("timestamp%d%d", timestamp, v), "values"
		if e := Put(key, val, nil); e != nil {
			log.Printf(" leveldb Put: Err\n [%s] %v\n", key, e)
		} else {
			log.Printf(" leveldb Put: Ok\n [%s] %s\n", key, val)
		}
	}
	// 遍历全部 keys
	RangeAll(nil, func(k, v []byte, ex error) {
		if ex != nil {
			log.Printf(" leveldb RangeAll: Err\n %v\n", ex)
		} else {
			log.Printf(" leveldb RangeAll: [%s] %s\n", string(k), string(v))
		}
	})
	// 遍历范围 keys: start ~ limit
	start, limit := fmt.Sprintf("timestamp%d%d", timestamp, 1), fmt.Sprintf("timestamp%d%d", time.Now().Unix(), 999)
	RangeWith(start, limit, nil, func(k, v []byte, ex error) {
		if ex != nil {
			log.Printf(" leveldb RangeWith: Err\n %v\n", ex)
		} else {
			log.Printf(" leveldb RangeWith: [%s] %s\n", string(k), string(v))
		}
	})
	// 遍历范围 keys: Start With prefix
	start = fmt.Sprintf("timestamp%d", timestamp)
	RangeStartWith(start, nil, func(k, v []byte, ex error) {
		if ex != nil {
			log.Printf(" leveldb RangeStartWith: Err\n %v\n", ex)
		} else {
			log.Printf(" leveldb RangeStartWith: [%s] %s\n", string(k), string(v))
		}
	})
}
