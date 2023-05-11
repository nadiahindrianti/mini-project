package database

import (
	"Mini_Project_Toko-Online/configs"
	"Mini_Project_Toko-Online/models"
)

func CreateProduct(p models.Product) (interface{}, error) {
	if err := configs.DB.Create(&p).Error; err != nil {
		return nil, err
	}

	if err := configs.DB.Joins("Category").Find(&p).Error; err != nil {
		return nil, err
	}
	return p, nil
}

func UpdateProductController(productID uint, p models.Product) (interface{}, error) {
	product := models.Product{}
	product.ID = productID
	if err := configs.DB.Joins("Category").Find(&product).Error; err != nil {
		return nil, err
	}

	product.Title = p.Title
	product.Price = p.Price
	product.Stock = p.Stock
	product.CategoryID = p.CategoryID

	err := configs.DB.Save(&product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

func GetProducts() (interface{}, error) {
	var products []models.Product

	if err := configs.DB.Joins("Category").Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func GetProduct(Category int) (interface{}, error) {
	var product []models.Product
	Category = int(Category)

	if err := configs.DB.Where("category_id = ?", Category).Preload("Category").Find(&product).Error; err != nil {
		return nil, err
	}

	return product, nil
}

func DeleteProductController(productID int) (interface{}, error) {
	var admin models.Admin
	if err := configs.DB.Delete(&admin, productID).Error; err != nil {
		return nil, err
	}
	return productID, nil
}
