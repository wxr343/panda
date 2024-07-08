package routev1

import (
	"github.com/gin-gonic/gin"
	"panda/app/controllers/app"
)

func SetUserGroupRoutes(router *gin.RouterGroup) {
	router.POST("/register", app.Register)
}
