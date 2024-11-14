package httputils

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
