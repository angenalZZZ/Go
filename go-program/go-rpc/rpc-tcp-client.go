package go_rpc

import (
	"context"
	"net/rpc"

	"github.com/angenalZZZ/Go/go-program/go-rpc/core"
)

type RpcTcpClient struct {
	client *rpc.Client
}

func (c *RpcTcpClient) Dial(addr string) (e error) {
	c.client, e = rpc.Dial("tcp", addr)
	return
}

func (c *RpcTcpClient) Close() (e error) {
	e = c.client.Close()
	return
}

func (c *RpcTcpClient) Execute(ctx context.Context, req *core.Request) (res *core.Response, e error) {
	res = new(core.Response)
	e = c.client.Call(core.HandlerName, req, res)
	return
}
