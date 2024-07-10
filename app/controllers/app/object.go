package app

import (
	"github.com/gin-gonic/gin"
	"panda/app/common/request"
	"panda/app/common/response"
	"panda/app/services"
)

// ObjectPush 添加物体
func ObjectPush(c *gin.Context) {
	var form request.Object
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}

	if err, object := services.ObjectService.ObjectPush(form); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, object)
	}
}

// ObjectQuery 查询物体
func ObjectQuery(c *gin.Context) {
	var form request.ObjectQuery
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}
	if err, object := services.ObjectService.ObjectQuery(form); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, object)
	}
}

// ObjectDelete 删除物体
func ObjectDelete(c *gin.Context) {
	var form request.ObjectDelete
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}
	if err := services.ObjectService.ObjectDelete(form); err != nil {
		e := "删除失败err：" + err.Error()
		response.BusinessFail(c, e)
	} else {
		info := form.Name + "删除成功"
		response.Success(c, info)
	}
}
