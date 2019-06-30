package go_rpc

import (
	"log"
	"net/http"

	"github.com/angenalZZZ/Go/go-program/go-rpc/proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type GrpcClient struct {
	conn   *grpc.ClientConn
	client proto.HandlerClient
}

func (c *GrpcClient) Dial(addr string) (e error) {
	if c.conn, e = grpc.Dial(addr, grpc.WithInsecure()); e == nil {
		c.client = proto.NewHandlerClient(c.conn)
	}
	return
}

func (c *GrpcClient) Close() (e error) {
	e = c.conn.Close()
	return
}

func (c *GrpcClient) Execute(ctx *gin.Context, request *proto.Request) (*proto.Response, error) {
	return c.client.Execute(ctx, request)
}

// 运行一个 GIN API 服务
func (c *GrpcClient) RunApi() {
	g := gin.Default()
	g.GET("/api/:a/:q", func(ctx *gin.Context) {
		a, q := ctx.Param("a"), ctx.Param("q")

		req := &proto.Request{Action: a, Query: q}
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
