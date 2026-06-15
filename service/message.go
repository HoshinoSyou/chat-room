package service

import (
	"chat-room/dao"
	"chat-room/models"
	"errors"
	"gorm.io/gorm"
	"log"
)

func CreateMessage(fromUser uint, toUsername string, content string) (bool, *models.Message, error) {
	user := dao.SelectUserByUsername(toUsername)
	if user.ID <= 0 {
		err := errors.New("未找到用户名：" + toUsername)
		log.Printf("来自用户 %d 发送给用户 %s 的消息发送失败！错误信息：%v", fromUser, toUsername, err)
		return false, nil, err
	}
	msg := &models.Message{
		Model:      gorm.Model{},
		FromUserId: fromUser,
		ToUserId:   user.ID,
		Content:    content,
	}
	err := dao.CreateMessage(msg)
	if err != nil {
		log.Printf("userId:%d 发送给 userId:%d 的消息发送失败！错误信息：%v", fromUser, user.ID, err)
		return false, nil, err
	}
	return true, msg, nil
}

func SelectMessages(user1 uint, user2 uint, limit int) (bool, []*models.Message, error) {
	messages, err := dao.SelectMessage(user1, user2, limit)
	if err != nil {
		log.Printf("查询 userId:%d 和 userId:%d 的消息记录失败！错误信息：%v", user1, user2, err)
		return false, nil, err
	}
	return true, messages, nil
}
