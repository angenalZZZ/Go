package go_tcp

import (
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
			log.Fatal(e)
		}
	} else {
		log.Fatal(e)
	}
}

// 后台运行 tcp Serve Shutdown
func TcpSvrShutdown() {
	log.Println("后台服务 tcp: Server exiting..")
	if tcpSvr != nil {
		if e := tcpSvr.Shutdown(nil); e != nil {
			log.Fatal(e)
		}
		tcpSvr = nil
	}
	//fmt.Println("-------------------------")
}
