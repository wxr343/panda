package app

import (
	"github.com/gin-gonic/gin"
	"panda/app/common/request"
	"panda/app/common/response"
	"panda/app/services"
)

func ObjectPush(c *gin.Context) {
	var form request.Object
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}

	if err, user := services.ObjectService.ObjectPush(form); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, user)
	}
}
