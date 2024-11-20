package httputils

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"server-zys/logs"
	"time"
)

var (
	defaultCode = -1
	defaultMsg  = "unknown error"
)

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

type CodeError interface {
	error
	Code() int
}

func SuccessWithData(data interface{}) interface{} {
	return Response{
		Code: 0,
		Msg:  "成功",
		Data: data,
	}
}

func Error(err error) interface{} {
	// 检查 err 是否实现了 CodeError 接口
	code := defaultCode
	msg := defaultMsg

	if err != nil {
		msg = err.Error()
		if codeErr, ok := err.(CodeError); ok {
			code = codeErr.Code()
		}
	}

	return Response{
		Code: code,
		Msg:  msg,
	}
}

func Get(ctx context.Context, url string, header, data map[string]string) ([]byte, error) {
	context.WithTimeout(ctx, 3*time.Second)
	client := resty.New()
	resp, err := client.R().
		SetHeaders(header).
		SetQueryParams(data).
		Get(url)
	if err != nil {
		logs.Error(ctx, fmt.Sprintf("Http Get Error, url: %s, err: %s", url, err.Error()))
		return nil, err
	}

	return resp.Body(), nil
}

func Post(ctx context.Context, url string, header map[string]string, data interface{}) ([]byte, error) {
	context.WithTimeout(ctx, 3*time.Second)
	client := resty.New()
	resp, err := client.R().
		SetHeaders(header).
		SetBody(data).
		Post(url)

	if err != nil {
		logs.Error(ctx, fmt.Sprintf("Http Post Error, url: %s, err: %s", url, err.Error()))
		return nil, err
	}
	return resp.Body(), nil
}
