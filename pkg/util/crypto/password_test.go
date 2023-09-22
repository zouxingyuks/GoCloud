package crypto

import (
	"testing"
)

func Test_password(t *testing.T) {
	testCases := []struct {
		plaintext string
		expected  bool
	}{
		// 正确密码
		{"Password123!", true},

		// 短密码
		{"Pass", true},
		{"12345", true},

		// 长密码: 最长 72 bytes
		{"ThisIsAnExcessivelyLongPasswordThatIsTo00000000000000000000000oLongToUse", true},

		// 含有无效字符
		{"P@sswórd", true},
		{"1*3$5&", true},

		// 弱密码
		{"password", true},
		{"123456", true},
		{"abcdef", true},

		// 哈希错误（示例哈希）
		{"HashError1", false},
		{"HashError2", false},

		// 盐错误
		{"CorrectPasswordWithWrongSalt", false},
		{"AnotherCorrectPasswordWithWrongSalt", false},
	}

	crypto := NewCrypto(PasswordCrypto)

	for _, tc := range testCases {
		t.Run(tc.plaintext, func(t *testing.T) {
			// 调用您的密码加密函数 EncryptPassword 并与 expected 进行比较
			ciphertext, err := crypto.Encrypt([]byte(tc.plaintext))
			if err != nil {
				t.Errorf("加密密码 %s 时出错：%v", tc.plaintext, err)
			}

			result, err := crypto.Check([]byte(tc.plaintext), []byte(ciphertext))
			if err != nil {
				t.Errorf("校验密码 %s 时出错：%v", tc.plaintext, err)
			}

			if result != tc.expected {
				t.Errorf("密码 %s 预期：%v，实际：%v", tc.plaintext, tc.expected, result)
			}
		})
	}
}
