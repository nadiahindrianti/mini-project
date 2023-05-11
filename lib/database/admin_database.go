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
	if err := configs.DB.Where("email = ? AND password = ?", admin.Email, admin.Password).First(&admin).Error; err != nil {
		return nil, err
	}

	return admin, nil
}

func GetAdminsControllerAll() (interface{}, error) {
	var admins []models.Admin

	if err := configs.DB.Find(&admins).Error; err != nil {
		return nil, err
	}
	return admins, nil
}

func GetAdminProfile(adminid int) (interface{}, error) {
	var admin []models.Admin

	if err := configs.DB.First(&admin).Error; err != nil {
		return nil, err
	}

	return admin, nil
}

func UpdateAdminController(adminID uint, a models.Admin) (interface{}, error) {
	admin := models.Admin{}
	admin.ID = adminID
	configs.DB.First(&admin)

	admin.Name = a.Name
	admin.Email = a.Email
	admin.Contact = a.Contact
	admin.Role = a.Role
	admin.Password = a.Password

	err := configs.DB.Save(&admin).Error
	if err != nil {
		return nil, err
	}
	return admin, nil
}

func DeleteAdminController(id any) (interface{}, error) {
	var admin models.Admin
	if err := configs.DB.Delete(&admin).Error; err != nil {
		return nil, err
	}

	return admin, nil
}

func GetUsersControllerAll() (interface{}, error) {
	var users []models.User

	if err := configs.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func GetUsersId(userid int) (interface{}, error) {
	var user []models.User

	if err := configs.DB.First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
