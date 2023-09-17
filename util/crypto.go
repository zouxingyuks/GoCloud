package util

const (
	PasswordCrypto = "password"
)

type Crypto interface {
	Encrypt(plaintext string) (ciphertext string)
	Decrypt(ciphertext string) (plaintext string)
}

func NewCrypto(object string) Crypto {
	switch object {
	case PasswordCrypto:
		return password{}
	default:
		return nil
	}
}

type password struct {
}

func (password) Encrypt(plaintext string) (ciphertext string) {
	//todo 密码的加密算法
	return plaintext
}
func (password) Decrypt(ciphertext string) (plaintext string) {
	//todo 密码的解密算法
	return plaintext
}
