package go_rpc

import (
	"log"
	"net"

	"github.com/golang/protobuf/proto"

	proto_fast "github.com/angenalZZZ/Go/go-program/go-rpc/proto-fast"
)

// 后台运行 tcp Serve Run
func DoGrpcFastSvrRun() {

	// 监听TCP服务
	l, e := net.Listen("tcp", grpcTcpAddr)
	if e != nil {
		log.Fatalf("Failed to run grpc server: %v", e)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			panic(err)
		}
		log.Println("new connect", conn.RemoteAddr())
		go readMessage(conn)
	}
}

func readMessage(conn net.Conn) *proto_fast.Response {
	b, r := make([]byte, 4096), &proto_fast.Response{}
	for {
		if i, e := conn.Read(b); e != nil {
			return nil
		} else if i > 10 {
			if e = proto.Unmarshal(b[:i], r); e != nil {
				return nil
			}
			break
		}
	}
	return r
}
