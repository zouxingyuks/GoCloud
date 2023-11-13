package conf

import (
	"GoCloud/pkg/email/model"
)

var defaultConfig = map[string]interface{}{
	"usercontroller": userController{
		EmailVerify: false,
	},
	"system": system{
		Mode:  "master",
		Debug: false,
		Host:  "localhost",
		Port:  "9090",
		Sessions: sessions{
			Store:  "memory",
			Secret: "your-session-secret",
		},
		HashIDSalt: "your-hash-id-salt",
	},
	"database": database{
		Type:        "UNSET",
		User:        "",
		Password:    "",
		Host:        "",
		Name:        "",
		TablePrefix: "",
		DBFile:      "cloudreve.db",
		Port:        3306,
		Charset:     "utf8",
		UnixSocket:  false,
	},
	"redis": redis{
		Network:  "tcp",
		Server:   "localhost:6379",
		User:     "",
		Password: "",
		DB:       0,
		PoolSize: 10,
	},
	"cors": cors{
		AllowOrigins:     []string{"UNSET"},
		AllowMethods:     []string{"PUT", "POST", "GET", "OPTIONS"},
		AllowHeaders:     []string{"Cookie", "X-Cr-Policy", "Authorization", "Content-Length", "Content-Type", "X-Cr-Path", "X-Cr-FileName"},
		AllowCredentials: false,
		ExposeHeaders:    nil,
		SameSite:         "default",
		Secure:           true,
	},
	"mail": model.Mail{
		Type: 0,
		Smtp: model.SMTP{
			Host:     "smtp.example.com",
			Port:     465,
			User:     "",
			Password: "",
			Name:     "",
			Address:  "",
		},
	},
	"site": site{
		Domain: "localhost",
	},
}
