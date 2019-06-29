package go_rpc

import (
	"context"
	"net/rpc"

	"github.com/angenalZZZ/Go/go-program/go-rpc/core"
)

type RpcHttpClient struct {
	client *rpc.Client
}

func (c *RpcHttpClient) Dial(addr string) (e error) {
	c.client, e = rpc.DialHTTP("tcp", addr)
	return
}

func (c *RpcHttpClient) Close() (e error) {
	e = c.client.Close()
	return
}

func (c *RpcHttpClient) Execute(ctx context.Context, req *core.Request) (res *core.Response, e error) {
	res = new(core.Response)
	e = c.client.Call(core.HandlerName, req, res)
	return
}
