package go_rpc

import (
	"io"
	"log"
	"net"
	"net/http"

	"github.com/golang/protobuf/proto"

	proto_fast "github.com/angenalZZZ/Go/go-program/go-rpc/proto-fast"

	"github.com/gin-gonic/gin"
)

type GrpcClientFast struct {
	conn net.Conn
}

func (c *GrpcClientFast) Dial(addr string) (e error) {
	c.conn, e = net.Dial("tcp", addr)
	return
}

func (c *GrpcClientFast) Close() (e error) {
	e = c.conn.Close()
	return
}

func (c *GrpcClientFast) Execute(ctx *gin.Context, request *proto_fast.Request) (*proto_fast.Response, error) {
	if b, e := proto.Marshal(request); e != nil {
		return nil, e
	} else if _, e = c.conn.Write(b); e != nil {
		return nil, e
	}
	return readResponseMessage(c.conn)
}

// 运行一个 GIN API 服务
func (c *GrpcClientFast) RunApi() {
	g := gin.Default()
	g.GET("/api/:a/:q", func(ctx *gin.Context) {
		a, q := ctx.Param("a"), ctx.Param("q")

		req := &proto_fast.Request{Action: a, Query: q}
		if res, e := c.Execute(ctx, req); e == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code":   res.Code,
				"result": res.Result,
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": e.Error()})
		}
	})
	if e := g.Run(":8080"); e != nil {
		log.Fatalf("Failed to run grpc client server: %v", e)
	}
}

func readResponseMessage(conn net.Conn) (r *proto_fast.Response, err error) {
	i, b := 0, make([]byte, 4096)
	r = &proto_fast.Response{}
	for {
		if i, err = conn.Read(b); i > 0 {
			err = proto.Unmarshal(b[:i], r)
			return
		} else if err != nil && err != io.EOF {
			return
		}
	}
}
