package routev1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func SetApiGroupRoutes(router *gin.RouterGroup) {
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "group")
	})
	router.GET("/shutdown", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "api Shutdown success")
	})
}
