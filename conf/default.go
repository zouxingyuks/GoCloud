package conf

import (
	"GoCloud/pkg/email/model"
	"github.com/gin-contrib/sessions"
	"net/http"
)

var defaultConfig = map[string]any{
	"static": static{
		Mode: "master",
		Host: "0.0.0.0",
		Port: "8281",
		Session: session{
			Store:  "memory",
			Secret: "your-session-secret",
			Option: sessions.Options{
				HttpOnly: true,
				//86400 是一天
				MaxAge:   1 * 86400,
				Path:     "/",
				SameSite: http.SameSiteDefaultMode,
				Secure:   false,
			},
		},
	},
	"service": service{User: userService{
		RegisterEnable: true,
		EmailVerify:    false,
		LoginLimit: struct {
			Enable bool
			Period int
			Count  int
		}{
			Enable: false,
			Period: 0,
			Count:  0,
		},
	}},
	"system": system{
		Debug:      false,
		HashIDSalt: "your-hash-id-salt",
	},
	"dao": dao{
		Database: database{
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
		SSL:    false,
	},
}
