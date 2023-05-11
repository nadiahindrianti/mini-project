package database

import (
	"Mini_Project_Toko-Online/configs"
	"Mini_Project_Toko-Online/models"
)

func LoginUser(user models.User) (interface{}, error) {
	var err error

	if err = configs.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil

}

func RegisterUser(user models.User) (interface{}, error) {
	err := configs.DB.Create(&user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

func UpdateUser(UserID uint, u models.User) (interface{}, error) {
	user := models.User{}
	user.ID = UserID
	configs.DB.First(&user)

	user.Name = u.Name
	user.Email = u.Email
	user.Contact = u.Contact
	user.Role = u.Role
	user.Password = u.Password

	err := configs.DB.Save(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetUser(userid int) (interface{}, error) {
	var user []models.User

	if err := configs.DB.First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func DeleteUser(id any) (interface{}, error) {
	var user models.User
	if err := configs.DB.Delete(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
