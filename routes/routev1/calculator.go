package routev1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"panda/app/controllers/app"
)

func SetCalculatorRoute(router *gin.RouterGroup) {
	router.POST("/pushObject", app.ObjectPush, func(c *gin.Context) {
		c.String(http.StatusOK, "Push")
	})

}
