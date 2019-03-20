package main

import (
	"angenalZZZ/go-program/api-config"
	"angenalZZZ/go-program/api-svr"
	"angenalZZZ/go-program/go-file"
	"angenalZZZ/go-program/go-leveldb"
	"angenalZZZ/go-program/go-opentsdb"
	"angenalZZZ/go-program/go-redis"
	"angenalZZZ/go-program/go-scheduler"
	"angenalZZZ/go-program/go-shutdown-hook"
	"angenalZZZ/go-program/go-ssdb"
	"angenalZZZ/go-program/go-tcp"
	"angenalZZZ/go-program/go-type"
	"flag"
	"time"
)

/**
命令行参数
*/
var config = flag.Bool("config", true, "check config file .env")

var typeCheck = flag.Bool("type-check", false, "test Type Check")
var createFile = flag.Bool("create-file", false, "test Create File")

var tcp = flag.Bool("tcp", false, "open tcp Serve")
var http = flag.Bool("http", true, "open http Serve")

var leveldb = flag.Bool("leveldb", false, "test leveldb Client")
var opentsdb = flag.Bool("opentsdb", false, "test opentsdb Client")
var redis = flag.Bool("redis", false, "test redis Client")
var redisCli = flag.Bool("redis-cli", false, "test redis Cli")
var ssdb = flag.Bool("ssdb", false, "test SSdb Client")

var worker = flag.Bool("worker", false, "test Scheduler Worker")

/**
程序初始化
*/
func init() {
	// log使用UTC时间
	//log.SetFlags(log.Ldate | log.Ltime | log.LUTC)

	// 查看命令行参数 -h -help
	flag.Parse()

	// 监听程序退出1 后台运行 tcp Serve Shutdown
	go_shutdown_hook.Add(go_tcp.TcpSvrShutdown)
	// 监听程序退出2 后台运行 http Serve Shutdown
	go_shutdown_hook.Add(api_svr.HttpSvrShutdown)
	// 监听程序退出3 数据库 Leveldb Client
	go_shutdown_hook.Add(go_leveldb.ShutdownClient)
	// 监听程序退出4 数据库 OpenTSDB Client
	go_shutdown_hook.Add(go_opentsdb.ShutdownClient)
	// 监听程序退出5 数据库 Redis Client
	go_shutdown_hook.Add(go_redis.ShutdownClient)
	go_shutdown_hook.Add(go_redis.ShutdownCli)
	// 监听程序退出6 数据库 SSdb Client
	go_shutdown_hook.Add(go_ssdb.ShutdownClient)
	// 监听程序退出7 计划任务 Scheduler Worker
	go_shutdown_hook.Add(go_scheduler.ShutdownWorker)

}

/**
程序入口函数
*/
func main() {
	defer shutdown()
	time.Sleep(time.Nanosecond * 10)

	// 加载配置文件并检查配置项
	if *config == true {
		api_config.LoadCheck()
	}

	// 类型检查
	if *typeCheck == true {
		go_type.TypeCheck()
	}

	// 文件管理：创建文件
	if *createFile == true {
		go_file.CreateFile()
	}

	// 内存数据库 Leveldb Client
	if *leveldb == true {
		go go_leveldb.Test()
	}
	// 时序数据库 OpenTSDB Client
	if *opentsdb == true {
		go go_opentsdb.Test()
	}
	// 缓存数据库 Redis Client
	if *redis == true {
		go go_redis.Test()
	}
	if *redisCli == true {
		go go_redis.TestCli()
	}
	// 缓存数据库 SSdb Client
	if *ssdb == true {
		go go_ssdb.Test()
	}

	// 计划任务 Scheduler Worker
	if *worker == true {
		go go_scheduler.TestWorker()
	}

	// 后台运行 tcp Serve Run
	if *tcp == true {
		go go_tcp.TcpSvrRun()
	}
	// 后台运行 http Serve Run
	if *http == true {
		go api_svr.HttpSvrRun()
	}
}

/**
程序退出, 正常时 os.Exit(0) | 异常时 os.Exit(1)
*/
func shutdown() {
	go_shutdown_hook.Wait()
}
