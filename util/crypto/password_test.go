package crypto

import "testing"

func Test_password(t *testing.T) {
	plaintext := []byte("123456")
	crypto := NewCrypto(PasswordCrypto)
	ciphertext, err := crypto.Encrypt(plaintext)
	if err != nil {
		t.Error(err)
	}
	result, err := crypto.Check(plaintext, []byte(ciphertext))
	if err != nil {
		t.Error(err)
	}
	if !result {
		t.Error("密码校验失败")
	}

}
