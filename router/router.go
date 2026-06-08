package router

import (
	"chat-room/config"
	"chat-room/controller"
	"chat-room/middleware"
	"github.com/gin-gonic/gin"
)

func Entrance() {
	router := gin.Default()
	user := router.Group("/user")
	{
		user.POST("/register", controller.Register)
		user.POST("/login", controller.Login)
		user.Use(middleware.CheckToken())
		user.POST("/update")
		user.POST("/delete")
	}
	router.Run(":" + config.AppConfig.Server.Port)
}
