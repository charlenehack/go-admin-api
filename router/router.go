package router

import (
	"admin-api/api/controller"
	"admin-api/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.Logger()) // 自定义日志中间件
	register(router)
	return router
}

func register(r *gin.Engine) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	api := r.Group("/api")
	api.POST("/v1/login", controller.Login)
	// 需要认证的v1路由
	v1Jwt := api.Group("/v1", middleware.AuthMiddleware())
	{
		v1Jwt.POST("/user/add", controller.CreateUser)
		v1Jwt.GET("/user/list", controller.GetUserList)
		v1Jwt.PUT("/user/updateStatus", controller.UpdateUserStatus)
	}
}
