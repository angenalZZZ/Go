package go_rpc

import (
	"context"
	"net/rpc"
	"net/rpc/jsonrpc"

	"github.com/angenalZZZ/Go/go-program/go-rpc/core"
)

type RpcJsonClient struct {
	client *rpc.Client
}

func (c *RpcJsonClient) Dial(addr string) (e error) {
	c.client, e = jsonrpc.Dial("tcp", addr)
	return
}

func (c *RpcJsonClient) Close() (e error) {
	e = c.client.Close()
	return
}

func (c *RpcJsonClient) Execute(ctx context.Context, req *core.Request) (res *core.Response, e error) {
	res = new(core.Response)
	e = c.client.Call(core.HandlerName, req, res)
	return
}
