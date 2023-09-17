package crypto

const (
	PasswordCrypto = iota
)

type Crypto interface {
	Encrypt(plaintext string) (ciphertext string)
	Decrypt(ciphertext string) (plaintext string)
}

func NewCrypto(object int) Crypto {
	switch object {
	case PasswordCrypto:
		return password{}
	default:
		return nil
	}
}
