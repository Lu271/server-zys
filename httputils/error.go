package httputils

type HttpError struct {
	code int
	msg  string
}

func (h HttpError) Error() string {
	return h.msg
}

func (h HttpError) Code() int {
	return h.code
}

func NewHttpError(code int, msg string) error {
	return HttpError{code: code, msg: msg}
}

var (
	UserNotLogin  = NewHttpError(1001, "user not login")
	InterNalError = NewHttpError(1002, "inter nal error")
)
