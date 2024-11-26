package middleware

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/utils"
	"time"
)

func LogMiddleware(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, request, response interface{}) (err error) {
		rpc := rpcinfo.GetRPCInfo(ctx)
		var req, resp interface{}
		if arg, ok := request.(utils.KitexArgs); ok {
			req = arg.GetFirstArgument()
		}
		start := time.Now()
		err = next(ctx, request, response)
		if arg, ok := response.(utils.KitexResult); ok {
			resp = arg.GetResult()
		}

		service := fmt.Sprintf("%v.%v", rpc.To().ServiceName(), rpc.To().Method())
		remoter := rpc.To().Address()
		duration := time.Now().Sub(start)
		logs := fmt.Sprintf("rpc=%v, remoter=%v duration=%v req=%v resp=%v",
			service, remoter, duration.Milliseconds(), req, resp)
		fmt.Println(logs)
		return
	}
}
