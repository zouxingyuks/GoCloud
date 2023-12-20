package active

import (
	"GoCloud/conf"
	"github.com/pkg/errors"
	"net/url"
)

func generateURL(uuid string) (string, error) {
	// 生成激活链接
	if token, err := Generate(activeAction, uuid, exp); err != nil {
		return "", errors.Wrap(err, "Error generating activation token")
	} else {
		u := url.URL{}
		u.Path, err = url.JoinPath("/users/activate/" + token)
		if err != nil {
			return "", errors.Wrap(err, "Error generating activation link")
		}
		u.Scheme = "http"
		u.Host = conf.SiteConfig().Domain
		if conf.SiteConfig().SSL {
			u.Scheme = "https"
		}
		return u.String(), nil
	}
}
