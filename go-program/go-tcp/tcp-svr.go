package go_tcp

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

/**
后台服务 tcp: Server
*/
var tcpSvr = &http.Server{Addr: ""}

// 后台运行 tcp Serve Run
func TcpSvrRun() {
	netListener, e := net.Listen("tcp4", tcpSvr.Addr)
	if e == nil {
		fmt.Printf("\n-------------------------\n%s 后台服务 tcp: Server starting..\n", time.Now().Format(time.RFC3339))
		log.Fatal(tcpSvr.Serve(netListener))
	} else {
		log.Fatal(e)
	}
}

// 后台运行 tcp Serve Shutdown
func TcpSvrShutdown() {
	fmt.Printf("-------------------------\n%s 后台服务 tcp: Server exiting..\n", time.Now().Format(time.RFC3339))
	if tcpSvr != nil {
		if e := tcpSvr.Shutdown(nil); e != nil {
			log.Fatal(e)
		}
		tcpSvr = nil
	}
	fmt.Printf("%s 退出应用 %s\n-------------------------\n", time.Now().Format(time.RFC3339), os.Args[0])
}
