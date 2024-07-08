package routev1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetHomeRoute(router *gin.RouterGroup) {
	router.GET("", func(c *gin.Context) {
		//c.Request.URL.Path = "/home"
		c.String(http.StatusOK, "home")
	})
}
