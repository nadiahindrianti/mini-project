package database

import (
	"Mini_Project_Toko-Online/configs"
	"Mini_Project_Toko-Online/models"
)

func RegisterAdmin(admin models.Admin) (interface{}, error) {
	err := configs.DB.Create(&admin).Error

	if err != nil {
		return nil, err
	}

	return admin, nil
}

func LoginAdmin(admin models.Admin) (interface{}, error) {
	var err error

	if err = configs.DB.Where("email = ? AND password = ?", admin.Email, admin.Password).First(&admin).Error; err != nil {
		return nil, err
	}
	return admin, nil
}

func GetAdminControllerAll() (interface{}, error) {
	var adminAll []models.Admin

	if err := configs.DB.Find(&adminAll).Error; err != nil {
		return nil, err
	}
	return adminAll, nil
}

func GetAdminProfile(adminID uint) (interface{}, error) {
	var admin models.Admin
	admin.ID = adminID

	if err := configs.DB.First(&admin).Error; err != nil {
		return nil, err
	}

	return admin, nil
}

func UpdateAdminController(admin models.Admin) (interface{}, error) {
	var adminUpdate models.Admin
	configs.DB.First(&adminUpdate, admin.ID)

	if e := configs.DB.Model(&adminUpdate).Updates(models.Admin{Name: admin.Name, Email: admin.Email, Role: admin.Role, Password: admin.Password}).Error; e != nil {
		return nil, e
	}
	return adminUpdate, nil

}

func DeleteAdminController(adminID int) (interface{}, error) {
	err := configs.DB.Delete(&models.Admin{}, adminID).Error
	if err != nil {
		return nil, err
	}

	return adminID, nil

}

func GetUsersControllerAll() (interface{}, error) {
	var userAll []models.User

	err := configs.DB.Find(&userAll).Error
	if err != nil {
		return nil, err
	}

	return userAll, nil
}

func GetUsersId(userID uint) (interface{}, error) {
	var user models.User
	user.ID = userID

	if err := configs.DB.First(&user).Error; err != nil {
		return nil, err
	}

	return userID, nil
}
