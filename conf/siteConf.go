package conf

import (
	"log"
	"sync"
)

type site struct {
	Domain string
	once   sync.Once
	SSL    bool
}

var siteInstance = new(site)

// SiteConfig 站点配置
func SiteConfig() *site {
	siteInstance.once.Do(
		func() {
			log.Println("inti siteConfig...")
			Config().Sub("site").Unmarshal(&siteInstance)
			log.Println("init siteConfig...end")
		})
	return siteInstance
}
