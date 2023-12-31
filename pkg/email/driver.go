package email

type Driver interface {
	// Submit 提交邮件，由协程池进行协调发送
	Submit(to, title, body string) error
	// send 发送邮件
	send(to, title, body string) error
	// Close 关闭驱动
	Close()
}

var driverInstance Driver

func NewDriver(driverType int) Driver {
	switch driverType {
	case SMTPType:
		return NewSMTP()
	}
	return nil
}
