package controller

import (
	"chat-room/service"
	"chat-room/util/response"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func CreateMessage(ctx *gin.Context) {
	uid, exists := ctx.Get("uid")
	if !exists {
		response.Error(ctx, "验证用户登录状态失败！", errors.New("验证用户登录状态失败！"))
		return
	}
	targetUser := ctx.PostForm("username")
	//targetUid, err := strconv.ParseUint(targetUser, 10, 32)
	//if err != nil {
	//	log.Printf("转换 userId 为 %s 的用户 ID 失败！错误信息：%v", targetUser, err)
	//	response.Error(ctx, "无效的用户 ID", err)
	//	return
	//}
	content := ctx.PostForm("content")
	res, _, err := service.CreateMessage(uid.(uint), targetUser, content)
	if !res {
		response.Error(ctx, "发送消息失败！", err)
		return
	}
	response.Ok(ctx, "发送消息成功！")
}

func SelectMessage(ctx *gin.Context) {
	uid, exists := ctx.Get("uid")
	if !exists {
		response.Error(ctx, "验证用户登录状态失败！", errors.New("验证用户登录状态失败！"))
		return
	}
	targetUser := ctx.Param("user_id")
	targetUid, err := strconv.ParseUint(targetUser, 10, 32)
	if err != nil {
		log.Printf("转换 userId 为 %s 的用户 ID 失败！错误信息：%v", targetUser, err)
		response.Error(ctx, "无效的用户 ID", err)
		return
	}
	limit := 50
	limitStr := ctx.Query("limit")
	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			log.Printf("limit string 转换为 int 类型失败！错误信息：%v", err)
			response.Error(ctx, "limit string 转换为 int 类型失败！", err)
			return
		}
	}
	res, messages, err := service.SelectMessages(uid.(uint), uint(targetUid), limit)
	if !res {
		response.Error(ctx, "获取聊天记录失败！", err)
		return
	}
	response.OkWithData(ctx, "获取聊天记录成功！", messages)
}
