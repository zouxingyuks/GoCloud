package conf

import (
	"fmt"
	"testing"
)

func TestAutoConf(t *testing.T) {
	AddPath("../../config")

	//fmt.Printf("SystemConfig:\n%+v\n", *SystemConfig())
	//fmt.Printf("UserConfig:\n%+v\n", *UserConfig())
	//fmt.Printf("DatabaseConfig:\n%+v\n", *DatabaseConfig())
	fmt.Printf("RedisConfig:\n%+v\n", *RedisConfig())
	fmt.Printf("CORSConfig:\n%+v\n", *CORSConfig())
	fmt.Printf("MailConfig:\n%+v\n", *MailConfig())
	fmt.Printf("SMTPConfig:\n%+v\n", *SMTPConfig())
}
