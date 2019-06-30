package go_rpc

import (
	"fmt"
	"log"
	"net"

	"github.com/angenalZZZ/Go/go-program/go-rpc/core"
	"github.com/angenalZZZ/Go/go-program/go-rpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

/**
后台服务 grpc-tcp: Server
*/
var grpcTcpAddr string

// 初始化配置
func init() {
	grpcTcpAddr = "127.0.0.1:8007" // 选填本地IP
}

// 后台运行 tcp Serve Run
func DoGrpcSvrRun() {
	// 功能
	handler := &core.GrpcHandler{
		Actions: map[string]func(req *proto.Request, res *proto.Response) (err error){},
	}

	handler.Actions["ping"] = func(req *proto.Request, res *proto.Response) (err error) {
		res.Code = 200
		res.Result = fmt.Sprintf("%s : %s", req.GetAction(), req.GetQuery())
		return
	}

	// 监听TCP服务
	l, e := net.Listen("tcp", grpcTcpAddr)
	if e != nil {
		log.Fatalf("Failed to run grpc server: %v", e)
	}
	// 注册RPC
	s := grpc.NewServer()
	proto.RegisterHandlerServer(s, handler)
	reflection.Register(s)
	if e := s.Serve(l); e != nil {
		log.Fatalf("Failed to run grpc server: %v", e)
	}
}
