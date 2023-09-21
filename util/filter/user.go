package filter

import "net/mail"

// filterFacade 门面模式的结构体，用于封装多种过滤操作
type filterFacade struct{}

var FilterFacade filterFacade

// IsValidEmail 检查给定的电子邮件地址是否合法
func (f *filterFacade) IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
