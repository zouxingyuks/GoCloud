package filter

import "net/mail"

// IsValidEmail 检查给定的电子邮件地址是否合法
func (f *filterFacade) IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
