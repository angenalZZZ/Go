package grpc_log

import (
	"context"
	"fmt"
	"github.com/angenalZZZ/gofunc/f"
	"github.com/angenalZZZ/gofunc/log"
	"github.com/angenalZZZ/gofunc/log/lager"
	"google.golang.org/grpc"
	"time"
)

var grpcChannel = make(chan string, 100)

func ClientInterceptor() grpc.UnaryClientInterceptor {

	go handleGrpcChannel()

	return func(ctx context.Context, method string,
		req, reply interface{}, cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {

		// 开始时间
		startTime := time.Now().UnixNano() / 1e6

		err := invoker(ctx, method, req, reply, cc, opts...)

		// 结束时间
		endTime := time.Now().UnixNano() / 1e6

		// 日志格式
		grpcLogMap := make(map[string]interface{})

		grpcLogMap["request_time"] = startTime
		grpcLogMap["request_data"] = req
		grpcLogMap["request_method"] = method

		grpcLogMap["response_data"] = reply
		grpcLogMap["response_error"] = err

		grpcLogMap["cost_time"] = fmt.Sprintf("%vms", endTime-startTime)

		grpcLogJson, _ := f.EncodeJson(grpcLogMap)

		grpcChannel <- string(grpcLogJson)

		return err
	}
}

func handleGrpcChannel() {
	for accessLog := range grpcChannel {
		log.Info("gRPC", lager.Data{
			"Log": accessLog,
		})
	}
	return
}
