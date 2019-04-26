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
var tcpConn = map[string]net.Conn{}
var tcpListener *net.Listener

// 初始化配置
func init() {
	// <多命令窗口> nc 127.0.0.1 8007 > 输入请求内容
	tcpAddr = ":8007" // 选填本地IP
}

// 后台运行 tcp Serve Run
func DoTcpSvrRun() {
	// 监听TCP服务
	l, e := net.Listen("tcp4", tcpAddr)
	if e == nil {
		println()
		tcpListener = &l
		log.Printf("后台服务 tcp: Server starting.. Addr: %s\n", l.Addr())

		for {
			// 阻塞等待用户发出连接请求后，继续等待下一个用户的连接请求
			c, e := l.Accept()
			if e != nil {
				fmt.Printf("  tcp accept error: %v\n", e)
				return
			}
			// 保存当前用户
			i := c.RemoteAddr().String()
			fmt.Printf("  tcp accept Addr: %s\n", i)
			tcpConn[i] = c

			// 接收用户发出的请求信息，一次缓冲1024字节l
			go func(c net.Conn, l int) {
				// 当前用户进入
				i := c.RemoteAddr().String()
				// 当前用户退出
				defer func(c net.Conn, i string) {
					delete(tcpConn, i)
					_ = c.Close()
				}(c, i)

				for {
					// 等待当前用户发出请求信息
					b := make([]byte, l)
					n, e := c.Read(b)
					if e != nil {
						fmt.Printf("  tcp read error: %v\n", e)
						return
					}

					// 输出当前用户请求信息给其他客户端
					h, s, x := make([]string, len(tcpConn)-1), []byte(fmt.Sprintf(" %s : %s", i, b[:n])), 0
					for k, f := range tcpConn {
						if i == f.RemoteAddr().String() {
							continue
						}
						_, e = f.Write(s)
						// 记录其他客户端连接异常
						if e != nil {
							h[x], x = k, x+1
							fmt.Printf("  tcp write error: %v\n", e)
						}
					}

					// 关闭其他客户端连接异常
					for _, j := range h {
						if j == "" {
							continue
						}
						if f, OK := tcpConn[j]; OK {
							_ = f.Close()
							delete(tcpConn, j)
						}
					}

					// 处理用户请求
				}
			}(c, 1024)
		}
	} else {
		log.Fatal(e) // 中断程序时输出
	}
}

// 后台运行 tcp Serve Shutdown
func ShutdownTcpSvr() {
	if tcpListener != nil {
		// 关闭客户端连接
		for _, f := range tcpConn {
			_ = f.Close()
		}
		// 关闭TCP服务
		log.Println("后台服务 tcp: Server exiting..")
		if e := (*tcpListener).Close(); e != nil {
			log.Fatal(e) // 中断程序时输出
		}
	}
}
