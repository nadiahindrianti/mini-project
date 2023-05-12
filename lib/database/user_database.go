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

func UpdateUserController(user models.User) (interface{}, error) {
	var userUpdate models.User
	configs.DB.First(&userUpdate, user.ID)

	if e := configs.DB.Model(&userUpdate).Updates(models.User{Name: user.Name, Email: user.Email, Contact: user.Contact, Alamat: user.Alamat, Role: user.Role, Password: user.Password}).Error; e != nil {
		return nil, e
	}
	return userUpdate, nil

}

func GetUserProfile(userID uint) (interface{}, error) {
	var user models.User
	user.ID = userID

	if err := configs.DB.First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func DeleteUser(userID int) (interface{}, error) {
	err := configs.DB.Delete(&models.User{}, userID).Error
	if err != nil {
		return nil, err
	}

	return userID, nil
}
