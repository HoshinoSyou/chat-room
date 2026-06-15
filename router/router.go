package router

import (
	"chat-room/config"
	"chat-room/controller"
	"chat-room/middleware"
	"github.com/gin-gonic/gin"
)

func Entrance() {
	router := gin.Default()
	router.StaticFile("/", "./static/index.html")
	// 用户相关路由组
	user := router.Group("/user")
	{
		user.POST("/register", controller.Register)
		user.POST("/login", controller.Login)
		user.Use(middleware.CheckToken())
		user.POST("/update")
		user.POST("/delete")
	}
	// 聊天相关路由组
	chat := router.Group("/chat")
	chat.Use(middleware.CheckToken())
	{
		chat.GET("/messages/:user_id", controller.SelectMessage)
		chat.POST("/messages", controller.CreateMessage)
	}
	// 升级为 ws 路由
	router.GET("/ws", middleware.WSCheckToken(), controller.InitWebsocket)
	router.Run(":" + config.AppConfig.Server.Port)
}
