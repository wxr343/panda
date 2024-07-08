package services

import (
	"errors"
	"panda/app/common/request"
	"panda/app/models"
	"panda/global"
)

type userService struct {
}

var UserService = new(userService)

// Register 用户注册
func (userService *userService) Register(params request.Register) (err error, user models.User) {
	var result = global.App.DB.Where("mobile = ?", params.Mobile).Select("id").First(&models.User{})
	if result.RowsAffected != 0 {
		err = errors.New("手机号已存在")
		return
	}
	user = models.User{
		Name:     params.Name,
		Mobile:   params.Mobile,
		Password: params.Password,
	}
	err = global.App.DB.Create(&user).Error
	return
}
