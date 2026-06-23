package dao

import (
	"chat-room/models"
)

func CreateUser(user *models.User) error {
	DB.AutoMigrate(&models.User{})

	err := DB.Model("users").Create(&user).Error
	//if err != nil {
	//	log.Printf("mysql 数据库新增数据失败！错误信息：%v", err)
	//	return err
	//}
	return err
}

func SelectUserById(id uint) (user models.User) {
	DB.Model(&user).Where("id", id).First(&user)
	//if user.ID <= 0 {
	//	err := errors.New("未查找到 ID！")
	//	log.Printf("mysql 数据库查找数据失败！错误信息：%v", err)
	//	return models.User{}, err
	//}
	return
}

func SelectUsersById(ids []uint) (users []models.UserInformation) {
	DB.Model(&models.User{}).Select("id,username,avatar").Where("id IN ?", ids).Find(&users)
	return
}

func SelectUserByUsername(username string) (user models.User) {
	DB.Model(&user).Where("username = ?", username).First(&user)
	//if u.ID != 0 {
	//	err := errors.New("用户名已存在！")
	//	log.Printf("新增用户名失败！错误原因：%v", err)
	//	return false, err
	//}
	return
}

func UpdateUser(user *models.User) error {
	err := DB.Model(&user).Updates(&user).Error
	//if err != nil {
	//	log.Printf("mysql 数据库更新数据失败！错误信息：%v", err)
	//	return err
	//}
	return err
}

func DeleteUser(user *models.User) error {
	err := DB.Model(&user).Delete(&user).Error
	//if err != nil {
	//	log.Printf("mysql 数据库删除数据失败！错误信息：%v", err)
	//	return err
	//}
	return err
}
