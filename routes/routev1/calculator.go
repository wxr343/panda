package routev1

import (
	"github.com/gin-gonic/gin"
	"panda/app/controllers/app"
)

func SetCalculatorRoute(router *gin.RouterGroup) {
	router.POST("/pushObject", app.ObjectPush)

	router.POST("/queryObject", app.ObjectQuery)

	router.POST("/deleteObject", app.ObjectDelete)
}
