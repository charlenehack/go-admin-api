// @title 通用后台管理系统
// @version 1.0
// @description 后台管理系统API接口文档
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
package main

import (
	"admin-api/common/config"
	_ "admin-api/docs"
	"admin-api/pkg/db"
	"admin-api/pkg/log"
	"admin-api/router"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func init() {
	// 初始化数据库连接
	db.InitDb()
}

func main() {
	// 加载日志
	log := log.Log()
	// 设置启动模式
	gin.SetMode(config.Config.Server.Model)
	// 初始化路由
	router := router.InitRouter()
	addr := config.Config.Server.Host + ":" + config.Config.Server.Port
	app := &http.Server{
		Addr:    addr,
		Handler: router,
	}
	// 启动服务
	go func() {
                log.Info("Starting http server on ", addr)
		if err := app.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Info("listen: ", err)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Info("Shutdown Server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := app.Shutdown(ctx); err != nil {
		log.Info("Server Shutdown:", err)
	}
	log.Info("Server exiting...")
}
