package bootstrap

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"panda/global"
	"panda/routes/routev1"
	"syscall"
	"time"
)

func setRouter() *gin.Engine {
	router := gin.Default()

	// 注册分组路由
	// homeGroup 首页
	homeGroup := router.Group("")
	routev1.SetHomeRoute(homeGroup)
	// apiGroup 部分api
	apiGroup := router.Group("/api")
	routev1.SetApiGroupRoutes(apiGroup)
	// userGroup 用户相关
	userGroup := router.Group("/user")
	routev1.SetUserGroupRoutes(userGroup)
	// calculator 计算器
	calculatorGroup := router.Group("/calculator")
	routev1.SetCalculatorRoute(calculatorGroup)

	return router
}

// RunServer 启动服务器
func RunServer() {
	r := setRouter()

	// 设置服务
	srv := &http.Server{
		Addr:    ":" + global.App.Config.App.Port,
		Handler: r,
	}
	// 协程启动
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			global.App.Log.Fatal("", zap.Any("listen info:", err))
		}
	}()
	// 等待退出信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	global.App.Log.Info("Shutdown Server ...")

	// 服务关闭前的处理工作
	ctx, chancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer chancel()

	if err := srv.Shutdown(ctx); err != nil {
		global.App.Log.Fatal("Server Shutdown:", zap.Any("err", err))
	}
	global.App.Log.Info("Server exiting")
	//r.Run(":" + global.App.Config.App.Port)
}
