package go_tcp

import (
	"fmt"
	"log"
	"net"
)

/**
后台服务 tcp: Server
*/
var tcpAddr string
var tcpConn []net.Conn
var tcpListener *net.Listener

// 初始化配置
func init() {
	// <多命令窗口> nc 127.0.0.1 8007 > 输入请求内容
	tcpAddr = "127.0.0.1:8007"
	//tcpConn = make([]net.Conn, 100)
}

// 后台运行 tcp Serve Run
func TestTcpSvrRun() {
	l, e := net.Listen("tcp4", tcpAddr)
	if e == nil {
		println()
		tcpListener = &l
		log.Printf("后台服务 tcp: Server starting.. Addr: %s\n", l.Addr())

		for {
			// 等待用户发出连接请求
			c, e := l.Accept()
			if e != nil {
				log.Printf("后台服务 tcp: Accept error: %v\n", e)
				return
			} else {
				tcpConn = append(tcpConn, c)
				log.Printf("后台服务 tcp: Accept Addr: %s\n", c.RemoteAddr())
			}

			// 接收用户发出的请求信息 一次最多可输入1024字节
			buf := make([]byte, 1024)
			n, e := c.Read(buf)
			if e != nil {
				fmt.Printf("后台服务 tcp: Read error: %v\n", e)
				return
			}
			for _, c1 := range tcpConn {
				buf = []byte(fmt.Sprintf(" %s : %s", c1.RemoteAddr(), buf[:n]))
				_, _ = c1.Write(buf)
			}

			// 处理和输出用户的请求
			//fmt.Printf("后台服务 tcp: Get Message\n    > %s\n", string(buf[:n]))
		}
	} else {
		log.Fatal(e) // 中断程序时输出
	}
}

// 后台运行 tcp Serve Shutdown
func TcpSvrShutdown() {
	if tcpListener != nil {
		log.Println("后台服务 tcp: Server exiting..")
		if e := (*tcpListener).Close(); e != nil {
			log.Fatal(e) // 中断程序时输出
		}
	}
}
