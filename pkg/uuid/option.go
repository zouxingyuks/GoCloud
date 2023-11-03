package uuid

type Option interface {
	apply(*option)
}

// funcOption 定义funcOption类型，实现 IOption 接口
type funcOption struct {
	f func(*option)
}

func (fo funcOption) apply(o *option) {
	fo.f(o)
}

func newFuncOption(f func(option *option)) Option {
	return &funcOption{
		f: f,
	}
}
