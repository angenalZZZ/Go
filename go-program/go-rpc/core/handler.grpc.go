package core

import (
	"context"
	"errors"

	"github.com/angenalZZZ/Go/go-program/go-rpc/proto"
)

type GrpcHandler struct {
	Actions map[string]func(req *proto.Request, res *proto.Response) (err error)
}

func (h *GrpcHandler) Execute(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	action := request.GetAction()
	if action == "" {
		return nil, errors.New("A action must be specified")
	}

	if handler, OK := h.Actions[action]; OK {
		res := &proto.Response{}
		err := handler(request, res)
		return res, err
	}

	return nil, errors.New("Action[" + action + "] is not found")
}
