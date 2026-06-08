package main

import (
	"chat-room/config"
	"chat-room/dao"
	"chat-room/router"
	"log"
)

func main() {
	err := config.InitConfig()
	if err != nil {
		log.Fatalf("配置文件加载失败！错误信息：%v", err)
		return
	}
	err = dao.InitMysql()
	if err != nil {
		log.Fatalf("MySQL 连接失败！错误信息：%v", err)
		return
	}
	err = dao.InitRedis()
	if err != nil {
		log.Fatalf("Redis 连接失败！错误信息：%v", err)
		return
	}
	router.Entrance()
}
