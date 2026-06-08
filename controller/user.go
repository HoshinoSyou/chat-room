package controller

import (
	"chat-room/models"
	"chat-room/service"
	"chat-room/util/response"
	"github.com/gin-gonic/gin"
	"log"
)

func Register(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		log.Printf("从请求获取 user 数据绑定参数失败！错误信息：%v", err)
		response.Error(ctx, "注册用户信息失败！", err)
		return
	}
	res, err := service.Register(&user)
	if !res {
		response.Error(ctx, "注册用户信息失败！", err)
		return
	}
	response.Ok(ctx, "注册用户信息成功！")
}

func Login(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	log.Println(user)
	if err != nil {
		log.Printf("从请求获取 user 数据绑定参数失败！错误信息：%v", err)
		response.Error(ctx, "用户登录失败！", err)
		return
	}
	res, err, token := service.Login(user.Username, user.Password)
	if !res {
		response.Error(ctx, "用户登录失败！", err)
		return
	}
	response.OkWithToken(ctx, "用户登录成功！", token)
}
