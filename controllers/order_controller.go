package controllers

import (
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
		"message": "success create new booking",
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
		"messages": "success delete product",
		"id":       result,
	})
}

func GetOrdersController(c echo.Context) error {
	bookings, err := database.GetOrders()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all",
		"bookings": bookings,
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

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get by id",
		"orders":   order,
	})
}
