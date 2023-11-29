package conf

import (
	"GoCloud/pkg/log"
	"github.com/gin-contrib/sessions"
	"github.com/pkg/errors"
	"sync"
)

type static struct {
	Mode    string
	Host    string
	Port    string
	Session session
}

type session struct {
	Store  string
	Secret string
	Option sessions.Options
}

var staticInstance = new(struct {
	once sync.Once
	*static
})

// StaticConfig 站点配置
func StaticConfig() *static {
	staticInstance.once.Do(
		func() {
			log.NewEntry("conf").Info("init siteConfig...start")
			err := Config().Sub("static").Unmarshal(&staticInstance.static)
			if err != nil {
				log.NewEntry("conf").Panic(errors.New("init siteConfig...failed").Error())
			}
			log.NewEntry("conf").Info("init siteConfig...done")
		})
	return staticInstance.static
}
