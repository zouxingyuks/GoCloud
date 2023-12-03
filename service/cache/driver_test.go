package cache

import (
	"fmt"
	"testing"
)

func testDriver() Driver {
	config := RedisConfig{
		Addr:     "100.76.246.116:6379",
		Password: "redis123",
		DB:       0,
		PoolSize: 10,
	}
	d, err := New(KindRedis, WithConf(config))
	if err != nil {
		panic(err)
	}
	return d
}
func TestSet(t *testing.T) {
	d := testDriver()

	err := d.Set("test", "test", 1000)
	if err != nil {
		panic(err)
	}

}
func TestHMSet(t *testing.T) {
	d := testDriver()
	err := d.HMSet("setting", map[string]interface{}{
		"siteName":                      "test",
		"login_captcha":                 "test",
		"reg_captcha":                   "test",
		"email_active":                  "test",
		"forget_captcha":                "test",
		"themes":                        "test",
		"defaultTheme":                  "test",
		"home_view_method":              "test",
		"share_view_method":             "test",
		"authn_enabled":                 "test",
		"captcha_ReCaptchaKey":          "test",
		"captcha_type":                  "test",
		"captcha_TCaptcha_CaptchaAppId": "test",
		"register_enabled":              "test",
		"show_app_promotion":            "test",
	}, 1000)
	if err != nil {
		panic(err)
	}
}

func TestHMGet(t *testing.T) {
	d := testDriver()
	// 正常获取
	val, err := d.HMGet("setting", []string{
		"siteName",
		"login_captcha",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(val)

	// 获取不存在的字段
	val, err = d.HMGet("setting", []string{
		"siteName",
		"login_captcha",
		"test",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(val)
}
