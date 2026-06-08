package dao

import (
	"chat-room/models"
)

func CreateUser(user *models.User) error {
	DB.AutoMigrate(&models.User{})

	err := DB.Table("users").Create(&user).Error
	//if err != nil {
	//	log.Printf("mysql 数据库新增数据失败！错误信息：%v", err)
	//	return err
	//}
	return err
}

func SelectUserById(id uint) (user models.User) {
	DB.Model(&models.User{}).Where("id", id).First(&user)
	//if user.ID <= 0 {
	//	err := errors.New("未查找到 ID！")
	//	log.Printf("mysql 数据库查找数据失败！错误信息：%v", err)
	//	return models.User{}, err
	//}
	return
}

func SelectUserByUsername(username string) (user models.User) {
	DB.Model(&models.User{}).Where("username", username).First(&user)
	//if u.ID != 0 {
	//	err := errors.New("用户名已存在！")
	//	log.Printf("新增用户名失败！错误原因：%v", err)
	//	return false, err
	//}
	return
}

func UpdateUser(user *models.User) error {
	err := DB.Model(&models.User{}).Updates(&user).Error
	//if err != nil {
	//	log.Printf("mysql 数据库更新数据失败！错误信息：%v", err)
	//	return err
	//}
	return err
}

func DeleteUser(user *models.User) error {
	err := DB.Model(&models.User{}).Delete(&user).Error
	//if err != nil {
	//	log.Printf("mysql 数据库删除数据失败！错误信息：%v", err)
	//	return err
	//}
	return err
}
