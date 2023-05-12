package controllers

import (
	"Mini_Project_Toko-Online/helpers"
	database "Mini_Project_Toko-Online/lib/database"
	"Mini_Project_Toko-Online/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateOrderController(c echo.Context) error {
	order := models.Order{}
	c.Bind(&order)

	result, err := database.CreateOrder(order)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Success Create Order",
		"order":   result,
	})

}

func DeleteOrderController(c echo.Context) error {
	orderID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	result, err := database.DeleteOrder(orderID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "Success Delete Order",
		"id":       result,
	})
}

func GetOrdersControllerAll(c echo.Context) error {
	order, err := database.GetOrders()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, helpers.ResponseData{
		Status:  "Success Get Order Data All",
		Message: "Successfuly",
		Data:    order,
	})
}

func GetOrderController(c echo.Context) error {
	User, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	order, err := database.GetOrderController(uint(User))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, helpers.ResponseData{
		Status:  "Success Get Order Data Id",
		Message: "Successfuly",
		Data:    order,
	})
}
