package main

import (
	"flag"
	"log"
	"math"
	"os"
	"runtime"
	"time"

	//_ "github.com/angenalZZZ/Go/go-program/api-config/env"
	_ "github.com/angenalZZZ/Go/go-program/api-config/env-viper"

	api_svr "github.com/angenalZZZ/Go/go-program/api-svr"
	go_file "github.com/angenalZZZ/Go/go-program/go-file"
	go_leveldb "github.com/angenalZZZ/Go/go-program/go-leveldb"
	go_opentsdb "github.com/angenalZZZ/Go/go-program/go-opentsdb"
	go_redis "github.com/angenalZZZ/Go/go-program/go-redis"
	go_scheduler "github.com/angenalZZZ/Go/go-program/go-scheduler"
	go_shutdown_hook "github.com/angenalZZZ/Go/go-program/go-shutdown-hook"
	go_ssdb "github.com/angenalZZZ/Go/go-program/go-ssdb"
	go_tcp "github.com/angenalZZZ/Go/go-program/go-tcp"
	go_type "github.com/angenalZZZ/Go/go-program/go-type"
)

/**
命令行参数
*/
var (
	flagTypeCheck  = flag.Bool("type-check", false, "test Type Check")
	flagCreateFile = flag.Bool("create-file", false, "test Create File")

	flagTcp  = flag.Bool("tcp", true, "open flagTcp Serve")
	flagHttp = flag.Bool("http", true, "open flagHttp Serve")

	flagLeveldb  = flag.Bool("leveldb", false, "test flagLeveldb Client")
	flagOpentsdb = flag.Bool("opentsdb", false, "test flagOpentsdb Client")
	flagRedis    = flag.Bool("redis", false, "test flagRedis Client")
	flagRedisCli = flag.Bool("redis-cli", false, "test flagRedis Cli")
	flagSsdb     = flag.Bool("ssdb", false, "test SSdb Client")

	flagWorker = flag.Bool("worker", false, "test Scheduler Worker")
)

/**
程序初始化
*/
func init() {

	// 命令行参数 查看 -h -help
	flag.Usage = func() {
		log.Printf(" Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	// 设置CPU空闲1个
	runtime.GOMAXPROCS(int(math.Max(float64(1), float64(runtime.NumCPU()-1))))

	// 监听程序退出
	// 1 flagTcp Serve
	go_shutdown_hook.Add(go_tcp.ShutdownTcpSvr)
	// 2 flagHttp Serve
	go_shutdown_hook.Add(api_svr.ShutdownHttpSvr)
	// 3 Leveldb Client
	go_shutdown_hook.Add(go_leveldb.ShutdownClient)
	// 4 OpenTSDB Client
	go_shutdown_hook.Add(go_opentsdb.ShutdownClient)
	// 5 Redis Client & Cli
	go_shutdown_hook.Add(go_redis.ShutdownClient)
	go_shutdown_hook.Add(go_redis.ShutdownCli)
	// 6 SSdb Client
	go_shutdown_hook.Add(go_ssdb.ShutdownClient)
	// 7 Scheduler Worker
	go_shutdown_hook.Add(go_scheduler.ShutdownWorker)

}

/**
程序入口函数
*/
func main() {
	defer end()

	//flag.Usage() // 查看程序使用说明
	time.Sleep(time.Nanosecond * 10)

	start()
}

/**
程序开始执行
*/
func start() {

	// 类型检查
	if *flagTypeCheck == true {
		go go_type.DoTypeCheck()
	}

	// 文件管理：创建文件
	if *flagCreateFile == true {
		go go_file.DoCreateFile()
	}

	// 内存数据库 Leveldb Client
	if *flagLeveldb == true {
		go go_leveldb.Do()
	}
	// 时序数据库 OpenTSDB Client
	if *flagOpentsdb == true {
		go go_opentsdb.Do()
	}
	// 缓存数据库 Redis Client
	if *flagRedis == true {
		go go_redis.Do()
	}
	if *flagRedisCli == true {
		go go_redis.DoCli()
	}
	// 缓存数据库 SSdb Client
	if *flagSsdb == true {
		go go_ssdb.Do()
	}

	// 计划任务 Scheduler Worker
	if *flagWorker == true {
		go go_scheduler.DoWorker()
	}

	// 后台运行 flagTcp Serve Run
	if *flagTcp == true {
		go go_tcp.DoTcpSvrRun()
	}
	// 后台运行 flagHttp Serve Run
	if *flagHttp == true {
		go api_svr.DoHttpSvrRun()
	}

}

/**
程序退出, 正常时 os.Exit(0) | 异常时 os.Exit(1)
*/
func end() {
	go_shutdown_hook.Wait()
}
