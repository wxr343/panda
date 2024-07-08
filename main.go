package main

import (
	logger "github.com/wxr343/logger/logger"
	"panda/bootstrap"
	"panda/global"
)

func main() {
	// 初始化配置
	bootstrap.InitializeConfig()
	// 初始化日志
	global.App.Log = logger.InitializeLog(global.LogCfg)
	//global.App.Log.Info("log init")

	// 初始化数据库
	global.App.DB = bootstrap.InitializeDB()

	// 初始化验证器
	bootstrap.InitializeValidator()

	//t := models.Statement{
	//	Model: gorm.Model{
	//		//ID:        1,
	//		CreatedAt: time.Now(),
	//		UpdatedAt: time.Now(),
	//		//DeletedAt: gorm.DeletedAt{},
	//	},
	//	Account:     "LPJ",
	//	Date:        time.Now(),
	//	Income:      11000,
	//	Expenditure: 1000,
	//}
	//bootstrap.InsertSqlData(global.App.DB, &t)

	// 程序关闭前，释放数据库连接
	defer func() {
		if global.App.DB != nil {
			db, _ := global.App.DB.DB()
			db.Close()
		}
	}()

	bootstrap.RunServer()
}
