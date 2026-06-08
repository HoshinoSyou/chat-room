package service

import (
	"chat-room/dao"
	"chat-room/models"
	"chat-room/util/jwt"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func Register(user *models.User) (bool, error) {
	u := dao.SelectUserByUsername(user.Username)
	if u.ID > 0 {
		err := errors.New("用户名已存在！")
		log.Printf("新增用户名失败！错误原因：%v", err)
		return false, err
	}

	bytesPwd, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("加密密码失败！错误信息：%v", err)
		return false, err
	}
	err = dao.CreateUser(&models.User{
		Username: user.Username,
		Password: string(bytesPwd),
		Nickname: user.Nickname,
		Avatar:   user.Avatar,
	})
	if err != nil {
		log.Printf("mysql 数据库新增数据失败！错误信息：%v", err)
		return false, err
	}
	return true, nil
}

func Login(username string, password string) (bool, error, string) {
	user := dao.SelectUserByUsername(username)
	//bytesPwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	//if err != nil {
	//	log.Printf("加密密码失败！错误信息：%v", err)
	//	return false, err
	//}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		err = errors.New("用户名密码错误！")
		log.Printf("用户名 %s 密码验证失败！", username)
		return false, err, ""
	}
	token, err := jwt.TokenCreate(username, user.ID)
	if err != nil {
		log.Printf("初始化登录状态失败！错误信息：%v", err)
		return false, err, ""
	}
	return true, nil, token
}
