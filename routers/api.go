package routers

import "github.com/gin-gonic/gin"

type api interface {
	load(r *gin.Engine)
}

func newApi(version int) api {
	switch version {
	case 0:
		return apiV0{}
	default:
		return apiV0{}
	}
}
