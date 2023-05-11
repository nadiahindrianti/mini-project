package database

import (
	"Mini_Project_Toko-Online/configs"
	"Mini_Project_Toko-Online/models"
)

func CreateCategoryController(category models.Category) (interface{}, error) {
	if err := configs.DB.Create(&category).Error; err != nil {
		return nil, err
	}

	if err := configs.DB.Joins("Category").Find(&category).Error; err != nil {
		return nil, err
	}
	return category, nil
}

func UpdateCategoryController(CategoryID uint, categoryupdate models.Category) (interface{}, error) {
	category := models.Category{}
	category.ID = CategoryID

	configs.DB.First(&category)

	category.Type = categoryupdate.Type
	category.Category = categoryupdate.Category

	err := configs.DB.Save(&category).Error
	if err != nil {
		return nil, err
	}
	return category, nil
}

func GetCategoryControllerAll() (interface{}, error) {
	var categories []models.Category

	if err := configs.DB.First(&categories).Error; err != nil {
		return nil, err
	}

	return categories, nil
}

func GetCategoryController(adminid int) (interface{}, error) {
	var category []models.Category

	if err := configs.DB.First(&category).Error; err != nil {
		return nil, err
	}

	return category, nil
}
