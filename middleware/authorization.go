package middleware

import (
	"chat-room/util/jwt"
	"chat-room/util/response"
	"encoding/base64"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
)

func CheckToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("Authorization")
		tokenSplit := strings.Split(token, ".")
		if len(tokenSplit) != 3 {
			err := errors.New("token 格式错误")
			log.Println(err)
			response.Error(ctx, "验证用户信息失败！", err)
			ctx.Abort()
			return
		}
		header, err := base64.StdEncoding.DecodeString(tokenSplit[0])
		if err != nil {
			log.Printf("token header 解析错误！错误信息：%v", err)
			response.Error(ctx, "验证用户信息失败！", err)
			ctx.Abort()
			return
		}
		headerSplit := strings.Split(string(header), " ")
		if !strings.EqualFold(headerSplit[0], "Bear ") {
			err := errors.New("token header 格式错误！")
			log.Println(err)
			response.Error(ctx, "验证用户信息失败！", err)
			ctx.Abort()
			return
		}

		payload, err := jwt.TokenParse(headerSplit[0] + tokenSplit[1] + tokenSplit[2])
		if err != nil {
			response.Error(ctx, "验证用户信息失败！", err)
			ctx.Abort()
			return
		}
		ctx.Set("uid", payload.Uid)
		ctx.Set("username", payload.Username)
		ctx.Next()
		//var payload jwt.Payload
		//err = json.Unmarshal([]byte(tokenSplit[1]), &payload)
		//if err != nil {
		//	log.Printf("unmarshal failed!反序列化失败！错误信息：%v", err)
		//	response.Error(ctx, "反序列化失败", err)
		//	ctx.Abort()
		//	return
		//}
	}
}
