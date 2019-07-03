package go_rpc

import (
	"context"
	"io"
	"log"
	"net"

	"github.com/golang/protobuf/proto"

	"github.com/angenalZZZ/Go/go-program/go-rpc/core"
	proto_fast "github.com/angenalZZZ/Go/go-program/go-rpc/proto-fast"
)

// 后台运行 tcp Serve Run
func DoGrpcFastSvrRun() {
	// 监听TCP服务
	l, e := net.Listen("tcp", grpcTcpAddr)
	if e != nil {
		log.Fatalf("Failed to run grpc server: %v", e)
	}

	// 功能
	handler := &core.GrpcFastHandler{
		Actions: map[string]func(req *proto_fast.Request, res *proto_fast.Response) (err error){},
	}

	handler.Actions["ping"] = func(req *proto_fast.Request, res *proto_fast.Response) (err error) {
		res.Code = 200
		res.Result = "pong"
		return
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			panic(err)
		}
		log.Println("new connect", conn.RemoteAddr())
		go writeResponseMessage(conn, handler)
	}
}

func readRequestMessage(conn net.Conn) (r *proto_fast.Request, err error) {
	i, b := 0, make([]byte, 4096)
	r = &proto_fast.Request{}
	for {
		if i, err = conn.Read(b); i > 0 {
			err = proto.Unmarshal(b[:i], r)
			return
		} else if err != nil && err != io.EOF {
			return
		}
	}
}

func writeResponseMessage(conn net.Conn, handler *core.GrpcFastHandler) *proto_fast.Response {
	if req, err := readRequestMessage(conn); err != nil {
		return nil
	} else if res, err := handler.Execute(context.Background(), req); err != nil {
		return nil
	} else {
		return res
	}
}
