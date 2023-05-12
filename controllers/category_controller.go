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
		Message: "Successfuly Create Category",
	})

}

func GetCategoryControllerAll(c echo.Context) error {
	var categoriAll []models.Category

	if err := configs.DB.Find(&categoriAll).Error; err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseNotData{
			Status:  "Error",
			Message: "Data Tidak Tersedia",
		})
	}

	return c.JSON(http.StatusOK, helpers.ResponseData{
		Status:  "Success Get Category All",
		Message: "Successfuly",
		Data:    categoriAll,
	})
}

func GetCategoryController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Invalid Id",
		})
	}

	var category models.Category
	if err = configs.DB.Where("id = ?", id).First(&category).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, helpers.ResponseData{
		Status:  "Success Get Category by ID",
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

	return c.JSON(http.StatusOK, helpers.ResponseData{
		Status:  "Success Update Category",
		Message: "Successfuly",
		Data:    category,
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
			"message": "Gagal Delete Data Category",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusBadRequest, helpers.ResponseNotData{
		Status:  "Succes Delete Data Category",
		Message: "Succesfuly",
	})
}
