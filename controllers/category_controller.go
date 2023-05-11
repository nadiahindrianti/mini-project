package controllers

import (
	"Mini_Project_Toko-Online/configs"
	"Mini_Project_Toko-Online/helpers"
	"Mini_Project_Toko-Online/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateCategoryController(c echo.Context) error {
	category := models.Category{}
	c.Bind(&category)

	if err := configs.DB.Save(&category).Error; err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseNotData{
			Status:  "Proses Create Category Gagal",
			Message: "Terdeteksi Eror",
		})
	}

	return c.JSON(http.StatusOK, helpers.ResponseNotData{
		Status:  "Proses Create Berhasil",
		Message: "Successfuly Registrasi Admin",
	})

}

func GetCategoryControllerAll(c echo.Context) error {
	var category []models.Category

	if err := configs.DB.Find(&category).Error; err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseNotData{
			Status:  "Error",
			Message: "Data Tidak Tersedia",
		})
	}

	return c.JSON(http.StatusOK, helpers.ResponseData{
		Status:  "Success Get Category",
		Message: "Successfuly",
		Data:    category,
	})
}

func UpdateCategoryController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseNotData{
			Status:  "Error",
			Message: "Data Tidak Tersedia",
		})
	}
	var category models.Category
	if err := configs.DB.Where("id = ?", id).First(&category).Error; err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseNotData{
			Status:  "Error",
			Message: "Data Tidak Tersedia",
		})
	}

	c.Bind(&category)
	if err := configs.DB.Model(&category).Updates(category).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, helpers.ResponseNotData{
		Status:  "Success Upload Data Category",
		Message: "Successfuly",
	})

}

func DeleteCategoryController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseNotData{
			Status:  "Error",
			Message: "Data Tidak Tersedia",
		})
	}

	var category models.Category
	if err := configs.DB.First(&category, "id = ? ", id).Error; err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseNotData{
			Status:  "Error",
			Message: "Data Tidak Tersedia",
		})
	}

	if err := configs.DB.Delete(&category).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Gagal Delete Category Data",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusBadRequest, helpers.ResponseNotData{
		Status:  "Succes Delete Data Category",
		Message: "Succesfuly",
	})
}
