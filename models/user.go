package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `form:"username" gorm:"uniqueIndex;size:32;not null" json:"username" binding:"required"`
	Password string `form:"password" gorm:"not null" json:"password" binding:"required"`
	Nickname string `form:"nickname" gorm:"size:50" json:"nickname"`
	Avatar   string `form:"avatar" gorm:"size:256" json:"avatar"`
}

type UserInformation struct {
	Id       uint   `json:"id" gorm:"column:id"`
	Username string `json:"username" gorm:"column:username"`
	Avatar   string `json:"avatar" gorm:"column:avatar"`
	//nickname string `json:"nickname"`
}
