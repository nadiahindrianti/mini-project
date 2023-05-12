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
	var product []models.Product

	if err := configs.DB.Joins("Category").Find(&product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func GetProductController(Category int) (interface{}, error) {
	var productAll []models.Product
	if err := configs.DB.Find(&productAll).Error; err != nil {
		return nil, err
	}

	return productAll, nil
}

func DeleteProductController(productID int) (interface{}, error) {
	err := configs.DB.Delete(&models.Product{}, productID).Error
	if err != nil {
		return nil, err
	}
	return productID, nil
}
