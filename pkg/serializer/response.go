package serializer

import (
	"GoCloud/pkg/conf"
	"GoCloud/pkg/log"
)

// Response 基础序列化器
type Response struct {
	Code  int         `json:"-"`
	Data  interface{} `json:"data,omitempty"`
	Msg   string      `json:"msg"`
	Error error       `json:"error,omitempty"`
}

func NewResponse(entry log.IEntry, Code int, opts ...ResponseOption) Response {
	o := new(responseOption)
	//在此处设置默认值
	for _, opt := range opts {
		opt.apply(o)
	}
	entry.Info(o.Msg, o.Fields...)
	o.Response.Code = Code
	return o.Response
}

type responseOption struct {
	Fields []log.Field
	Response
}

// ResponseOption 定义一个接口类型
type ResponseOption interface {
	apply(*responseOption)
}

// funcOption 定义funcOption类型，实现 IOption 接口
type funcOption struct {
	f func(*responseOption)
}

func newFuncOption(f func(option *responseOption)) ResponseOption {
	return &funcOption{
		f: f,
	}
}
func (fo funcOption) apply(o *responseOption) {
	fo.f(o)
}

// WithData 定义一个函数，用于设置 Data
func WithData(data interface{}) ResponseOption {
	return newFuncOption(func(o *responseOption) {
		o.Data = data
	})
}

// WithField 定义一个函数，用于设置 Fields
func WithField(fields ...log.Field) ResponseOption {
	return newFuncOption(func(o *responseOption) {
		o.Fields = append(o.Fields, fields...)
	})
}

// WithErr 将应用error携带标准库中的error
func WithErr(err error) ResponseOption {
	return newFuncOption(func(o *responseOption) {
		// 生产环境隐藏底层报错
		if err != nil && conf.SystemConfig().Debug {
			o.Error = err
			o.Fields = append(o.Fields, log.Field{
				Key:   "Error",
				Value: err,
			})
		}
	})
}

func WithMsg(Msg string) ResponseOption {
	return newFuncOption(func(o *responseOption) {
		o.Msg = Msg
	})
}
