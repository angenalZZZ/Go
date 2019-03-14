package main

import (
	"angenalZZZ/go-program/api-config"
	"angenalZZZ/go-program/go-file"
	"angenalZZZ/go-program/go-leveldb"
	"angenalZZZ/go-program/go-opentsdb"
	"angenalZZZ/go-program/go-redis"
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
var http = flag.Bool("http", false, "open http Serve")

var leveldb = flag.Bool("leveldb", false, "test leveldb Client")
var opentsdb = flag.Bool("opentsdb", false, "test opentsdb Client")
var redis = flag.Bool("redis", false, "test redis Client")
var ssdb = flag.Bool("ssdb", false, "test SSdb Client")

/**
程序入口函数
*/
func main() {
	// 查看命令行参数 -h -help
	flag.Parse()
	time.Sleep(time.Nanosecond * 100)

	// 监听程序退出1 后台运行 tcp Serve Shutdown
	if *tcp == true {
		go_shutdown_hook.Add(go_tcp.TcpSvrShutdown)
	}
	// 监听程序退出2 后台运行 http Serve Shutdown
	if *http == true {
		go_shutdown_hook.Add(go_tcp.HttpSvrShutdown)
	}
	// 监听程序退出3 数据库 Leveldb Client
	if *leveldb == true {
		go_shutdown_hook.Add(go_leveldb.ShutdownClient)
	}
	// 监听程序退出4 数据库 OpenTSDB Client
	if *opentsdb == true {
		go_shutdown_hook.Add(go_opentsdb.ShutdownClient)
	}
	// 监听程序退出5 数据库 Redis Client
	if *redis == true {
		go_shutdown_hook.Add(go_redis.ShutdownClient)
	}
	// 监听程序退出6 数据库 SSdb Client
	if *ssdb == true {
		go_shutdown_hook.Add(go_ssdb.ShutdownClient)
	}

	// 类型检查
	if *typeCheck == true {
		go_type.TypeCheck()
	}

	// 命令行参数
	//go_args.ArgsCheck()

	// 加载配置文件并检查配置项
	if *config == true {
		api_config.LoadCheck()
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
	// 缓存数据库 SSdb Client
	if *ssdb == true {
		go go_ssdb.Test()
	}

	// 后台运行 tcp Serve Run
	if *tcp == true {
		go go_tcp.TcpSvrRun()
	}
	// 后台运行 http Serve Run
	if *http == true {
		go go_tcp.HttpSvrRun()
	}

	// 程序退出时 os.Exit(0)
	go_shutdown_hook.Wait()
}
