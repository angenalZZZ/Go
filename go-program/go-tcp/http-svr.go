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
后台服务 http: Server
*/
var httpSvr = &http.Server{Addr: ":8008"}

// 后台运行 http Serve Run
func HttpSvrRun() {

	// 服务处理
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, e := fmt.Fprintf(writer, "%#v\n", request.URL)
		log.Fatal(e)
	})

	// 开始服务
	//log.Fatal(http.ListenAndServe(httpSvr.Addr, nil))
	netListener, e := net.Listen("tcp4", httpSvr.Addr)
	if e == nil {
		fmt.Printf("\n-------------------------\n%s 后台服务 http: Server starting.. host< %s >\n", time.Now().Format(time.RFC3339), httpSvr.Addr)
		log.Fatal(httpSvr.Serve(netListener))
	} else {
		log.Fatal(e)
	}
}

// 后台运行 http Serve Shutdown
func HttpSvrShutdown() {
	fmt.Printf("-------------------------\n%s 后台服务 http: Server exiting..\n", time.Now().Format(time.RFC3339))
	if httpSvr != nil {
		if e := httpSvr.Shutdown(nil); e != nil {
			log.Fatal(e)
		}
		httpSvr = nil
	}
	fmt.Printf("%s 退出应用 %s\n-------------------------\n", time.Now().Format(time.RFC3339), os.Args[0])
}
