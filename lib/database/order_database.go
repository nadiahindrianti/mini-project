package database

import (
	"Mini_Project_Toko-Online/configs"
	"Mini_Project_Toko-Online/models"
)

func CreateOrder(order models.Order) (interface{}, error) {
	err := configs.DB.Create(&order).Error

	if err != nil {
		return nil, err
	}

	return order, nil
}

func DeleteOrder(orderID int) (interface{}, error) {
	err := configs.DB.Delete(&models.Order{}, orderID).Error

	if err != nil {
		return nil, err
	}
	return orderID, nil
}

func GetOrders() (interface{}, error) {
	var orders []models.Order

	if err := configs.DB.Joins("Product").Joins("User").Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func GetOrderController(OrderID uint) (interface{}, error) {
	var order models.Order
	order.ID = OrderID

	if err := configs.DB.Preload("Product").Preload("User").Find(&order).Error; err != nil {
		return nil, err
	}

	return order, nil
}
