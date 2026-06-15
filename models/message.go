package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	FromUserId uint   `gorm:"index; not null" json:"from_user_id"`
	ToUserId   uint   `gorm:"index; not null" json:"to_user_id"`
	Content    string `gorm:"type:text" json:"content"`
	IsRead     bool   `gorm:"default:false" json:"is_read"`
}
