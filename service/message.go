package service

import (
	"chat-room/dao"
	"chat-room/models"
	"context"
	"encoding/json"
	"errors"
	"fmt"
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
	messages, err := dao.SelectHistoryMessages(user1, user2, limit)
	if err != nil {
		log.Printf("查询 userId:%d 和 userId:%d 的消息记录失败！错误信息：%v", user1, user2, err)
		return false, nil, err
	}
	return true, messages, nil
}

func CreateOfflineMessage(ctx context.Context, message *models.Message) (bool, error) {
	msg, err := json.Marshal(message)
	if err != nil {
		log.Printf("序列化消息失败！错误信息：%v", err)
		return false, err
	}
	key := fmt.Sprintf("offline_message:user_id:%d", message.ToUserId)
	err = dao.Rdb.LPush(ctx, key, msg).Err()
	if err != nil {
		log.Printf("存储 Redis 数据失败！错误信息：%v", err)
		return false, err
	}
	return true, nil
}

func SelectOfflineMessage(ctx context.Context, userId uint) (bool, []models.Message, error) {
	key := fmt.Sprintf("offline_message:user_id:%d", userId)
	length := dao.Rdb.LLen(ctx, key).Val()
	if length == 0 {
		return true, nil, nil
	}

	var messages []models.Message
	for i := int64(0); i < length; i++ {
		data, err := dao.Rdb.RPop(ctx, key).Bytes()
		if err != nil {
			log.Printf("从 Redis 中获取消息失败！错误信息：%v", err)
			return false, nil, err
		}

		var message models.Message
		err = json.Unmarshal(data, &message)
		if err != nil {
			log.Printf("反序列化失败！错误信息：%v", err)
			return false, nil, err
		}
		messages = append(messages, message)
	}
	return true, messages, nil
}

func SelectConversations(userId uint) (bool, []map[string]interface{}, error) {
	tempMap := make(map[uint]models.Message)
	messages, err := dao.SelectMessagesDESC(userId)
	if err != nil {
		log.Printf("查询用户 ID 为 %d 的聊天消息失败！错误信息：%v", userId, err)
		return false, nil, err
	}
	// 去重，后续考虑在 SQL 语句上优化去重步骤
	// TODO:用 SQL 语句进行去重步骤
	var friendIds []uint
	for _, msg := range messages {
		var friendId uint
		if msg.FromUserId == userId {
			friendId = msg.ToUserId
		} else {
			friendId = msg.FromUserId
		}
		_, exists := tempMap[friendId]
		if !exists {
			tempMap[friendId] = *msg
			friendIds = append(friendIds, friendId)
		}
	}

	// 通过 userId 批量查询 username
	users := dao.SelectUsersById(friendIds)
	length := len(users)
	userMap := make(map[uint]string, length)
	for _, user := range users {
		userMap[user.Id] = user.Username
	}

	// 组装成响应数据
	conversations := make([]map[string]interface{}, length)
	for friendId, conversation := range tempMap {
		conversations = append(conversations, map[string]interface{}{
			"friend_id":       friendId,
			"friend_username": userMap[friendId],
			"message_id":      conversation.ID,
			"time":            conversation.CreatedAt,
			"content":         conversation.Content,
		})
	}

	return true, conversations, nil
}
