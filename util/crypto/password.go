package crypto

import (
	"golang.org/x/crypto/bcrypt"
)

type password struct {
}

// Encrypt 加密
// 使用 bcrypt 加密
func (password) Encrypt(plaintext []byte) (string, error) {
	// 生成哈希密码
	ciphertext, err := bcrypt.GenerateFromPassword(plaintext, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(ciphertext), nil
}
func (password) Decrypt(ciphertext []byte) (plaintext string, err error) {
	return "", nil
}

func (p password) Check(plaintext []byte, ciphertext []byte) (result bool, err error) {
	err = bcrypt.CompareHashAndPassword(ciphertext, plaintext)
	if err != nil {
		return false, err
	}
	return true, nil
}
