package go_tcp

import (
	"context"
	"log"
	"net"
	"net/http"
)

/**
后台服务 tcp: Server
*/
var tcpSvr = &http.Server{Addr: ":8007"}

// 后台运行 tcp Serve Run
func TcpSvrRun() {
	l, e := net.Listen("tcp4", tcpSvr.Addr)
	if e == nil {
		println()
		log.Printf("后台服务 tcp: Server starting.. Addr: %s\n", tcpSvr.Addr)
		if e = tcpSvr.Serve(l); e != nil {
			log.Fatal(e) // 中断程序时输出
		}
	} else {
		log.Fatal(e) // 中断程序时输出
	}
}

// 后台运行 tcp Serve Shutdown
func TcpSvrShutdown() {
	log.Println("后台服务 tcp: Server exiting..")
	if tcpSvr != nil {
		if e := tcpSvr.Shutdown(context.Background()); e != nil {
			log.Fatal(e) // 中断程序时输出
		}
		tcpSvr = nil
	}
}
