package main

import (
	"angenalZZZ/go-program/api-config"
	"angenalZZZ/go-program/go-args"
	"angenalZZZ/go-program/go-file"
	"angenalZZZ/go-program/go-shutdown-hook"
	"angenalZZZ/go-program/go-tcp"
	"angenalZZZ/go-program/go-type"
)

/**
程序入口函数
*/
func main() {
	// 监听程序退出1 后台运行 tcp Serve Shutdown
	go_shutdown_hook.Add(go_tcp.TcpSvrShutdown)
	// 监听程序退出2 后台运行 http Serve Shutdown
	go_shutdown_hook.Add(go_tcp.HttpSvrShutdown)

	// 类型检查
	go_type.TypeCheck()

	// 命令行参数
	go_args.ArgsCheck()

	// 加载配置文件并检查配置项
	api_config.LoadCheck()

	// 文件管理：创建文件
	go_file.CreateFile()

	// 后台运行 tcp Serve Run
	go go_tcp.TcpSvrRun()
	// 后台运行 http Serve Run
	go go_tcp.HttpSvrRun()

	// 程序退出时
	go_shutdown_hook.Wait()
}
