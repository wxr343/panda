package utils

import (
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"panda/global"
)

// 密码加密及验证密码的方法

// BcryptMake 加密
func BcryptMake(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		global.App.Log.Info("BcryptMake", zap.Any("err", err))
	}
	return string(hash)
}

// BcryptMakeCheck 解密
func BcryptMakeCheck(pwd []byte, hashedPwd string) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, pwd)
	if err != nil {
		return false
	}
	return true
}
