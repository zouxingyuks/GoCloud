package conf

import (
	"GoCloud/pkg/log"
	"sync"
)

type site struct {
	Domain string
	once   sync.Once
}

var siteInstance = new(site)

// SiteConfig 站点配置
func SiteConfig() *site {
	siteInstance.once.Do(
		func() {
			log.NewEntry("inti siteConfig...")
			Config().Sub("site").Unmarshal(&siteInstance)
			log.NewEntry("init siteConfig...end")
		})
	return siteInstance
}
