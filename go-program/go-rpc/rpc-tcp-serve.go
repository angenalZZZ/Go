package go_rpc

import (
	"net"
	"net/rpc"

	"github.com/angenalZZZ/Go/go-program/go-rpc/core"
)

/**
后台服务 rpc-tcp: Server
*/
var rpcTcpAddr string

// 初始化配置
func init() {
	// <多命令窗口> nc 127.0.0.1 8007 > 输入请求内容
	rpcTcpAddr = "127.0.0.1:8007" // 选填本地IP
}

// 后台运行 tcp Serve Run
func DoRPCTcpSvrRun() {
	// 功能
	handler := &core.Handler{
		Actions: map[string]func(req *core.Request, res *core.Response) (err error){},
	}

	handler.Actions["ping"] = func(req *core.Request, res *core.Response) (err error) {
		res.Code = 200
		res.Result = "pong"
		return
	}

	// 注册RPC
	e := rpc.Register(handler)
	if e != nil {
		panic(e)
	}
	// 监听TCP服务
	l, e := net.Listen("tcp", rpcHttpAddr)
	if e != nil {
		panic(e)
	}
	rpc.Accept(l)
}
