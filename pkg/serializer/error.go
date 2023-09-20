package serializer

import (
	"errors"
	"github.com/gin-gonic/gin"
)

// AppError 应用错误
// 实现了error接口，可以自定义错误信息
type AppError struct {
	Code     int
	Msg      string
	RawError error
}

// NewError 返回新的错误对象
func NewError(code int, msg string, err error) AppError {
	return AppError{
		Code:     code,
		Msg:      msg,
		RawError: err,
	}
}

// NewErrorFromResponse 从 serializer.Response 构建错误
func NewErrorFromResponse(resp *Response) AppError {
	return AppError{
		Code:     resp.Code,
		Msg:      resp.Msg,
		RawError: errors.New(resp.Error),
	}
}

// WithError 将应用error携带标准库中的error
func (err *AppError) WithError(raw error) AppError {
	err.RawError = raw
	return *err
}

// Error 返回业务代码确定的可读错误信息
func (err AppError) Error() string {
	return err.Msg
}

// Err 通用错误处理
func Err(errCode int, msg string, err error) Response {
	var appError AppError
	// todo 解析这部分代码
	// 如果err是AppError类型，则尝试从中获取详细信息，否则这是自定义错误
	if errors.As(err, &appError) {
		errCode = appError.Code
		err = appError.RawError
		msg = appError.Msg
	}
	res := Response{
		Code: errCode,
		Msg:  msg,
	}
	// 生产环境隐藏底层报错
	if err != nil && gin.Mode() != gin.ReleaseMode {
		res.Error = err.Error()
	}
	return res
}
