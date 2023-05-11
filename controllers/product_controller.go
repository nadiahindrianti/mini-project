package controllers

import (
	"Mini_Project_Toko-Online/configs"
	"Mini_Project_Toko-Online/helpers"
	database "Mini_Project_Toko-Online/lib/database"
	"Mini_Project_Toko-Online/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateProductController(c echo.Context) error {
	product := models.Product{}
	c.Bind(&product)

	if err := configs.DB.Save(&product).Error; err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseNotData{
			Status:  "Proses Create Product Gagal",
			Message: "Terdeteksi Eror",
		})
	}

	return c.JSON(http.StatusOK, helpers.ResponseNotData{
		Status:  "Proses Create Berhasil",
		Message: "Successfuly Registrasi Admin",
	})

}

func GetProductControllerAll(c echo.Context) error {
	products, err := database.GetProducts()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all products",
		"products": products,
	})
}

func GetProductsController(c echo.Context) error {
	Category, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	product, err := database.GetProduct(Category)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get product",
		"products": product,
	})
}

func UpdateProductController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseNotData{
			Status:  "Error",
			Message: "Data Tidak Tersedia",
		})
	}
	var product models.Product
	if err := configs.DB.Where("id = ?", id).First(&product).Error; err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseNotData{
			Status:  "Error",
			Message: "Data Tidak Tersedia",
		})
	}

	c.Bind(&product)
	if err := configs.DB.Model(&product).Updates(product).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, helpers.ResponseNotData{
		Status:  "Success Upload Data Product",
		Message: "Successfuly",
	})

}

func DeleteProductController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseNotData{
			Status:  "Error",
			Message: "Data Tidak Tersedia",
		})
	}

	var product models.Product
	if err := configs.DB.First(&product, "id = ? ", id).Error; err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseNotData{
			Status:  "Error",
			Message: "Data Tidak Tersedia",
		})
	}

	if err := configs.DB.Delete(&product).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Gagal Delete User Data",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusBadRequest, helpers.ResponseNotData{
		Status:  "Succes Delete Data Product",
		Message: "Succesfuly",
	})
}
