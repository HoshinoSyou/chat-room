package dao

import "chat-room/models"

func CreateMessage(message *models.Message) error {
	DB.AutoMigrate(&models.Message{})
	err := DB.Model(&message).Create(&message).Error
	return err
}

func SelectMessage(user1 uint, user2 uint, limit int) ([]*models.Message, error) {
	var messages []*models.Message
	err := DB.Model(&models.Message{}).
		Where("(from_user_id = ? AND to_user_id = ?) OR (from_user_id = ? AND to_user_id = ?)", user1, user2, user2, user1).
		Order("created_at ASC").Limit(limit).Find(&messages).Error
	return messages, err
}
