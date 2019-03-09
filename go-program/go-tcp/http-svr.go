package go_tcp

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

/**
后台服务 http: Server
*/
var httpSvr *http.Server

// 后台运行 http Serve Run
func HttpSvrRun() {
	httpSvr = &http.Server{Addr: os.Getenv("HOST") + ":" + os.Getenv("POST")}

	// 服务处理
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {

		// 跟踪请求
		log.Printf("后台服务 http: %v\n", request.URL.String())

		// 处理请求
		_, e := fmt.Fprintf(writer, " %v %+v \n", time.Now(), request.URL)

		// 跟踪异常
		if e != nil {
			log.Fatal(e)
		}
	})

	// 开始服务
	//log.Fatal(http.ListenAndServe(httpSvr.Addr, nil))
	l, e := net.Listen("tcp4", httpSvr.Addr)
	if e == nil {
		println()
		log.Printf("后台服务 http: Server starting.. Addr: %s\n", httpSvr.Addr)
		if e = httpSvr.Serve(l); e != nil {
			log.Fatal(e)
		}
	} else {
		log.Fatal(e)
	}
}

// 后台运行 http Serve Shutdown
func HttpSvrShutdown() {
	if httpSvr != nil {
		log.Println("后台服务 http: Server stopping..") // Go ^1.8
		if e := httpSvr.Shutdown(context.Background()); e != nil {
			log.Fatal(e)
		}
	}
}