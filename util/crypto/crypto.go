package crypto

const (
	PasswordCrypto = iota
)

// Crypto
// 选用 []byte 主要是 此处无需进行字符串比较,且字符串操作较多，使用 []byte 效率更高
type Crypto interface {
	Encrypt(plaintext []byte) (ciphertext string, err error)
	Decrypt(ciphertext []byte) (plaintext string, err error)
}

func NewCrypto(object int) Crypto {
	switch object {
	case PasswordCrypto:
		return password{}
	default:
		return nil
	}
}
