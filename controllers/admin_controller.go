package controllers

import (
	"Mini_Project_Toko-Online/configs"
	"Mini_Project_Toko-Online/helpers"
	database "Mini_Project_Toko-Online/lib/database"
	"Mini_Project_Toko-Online/middlewares"
	"Mini_Project_Toko-Online/models"
	"net/http"
	"reflect"
	"strconv"

	"github.com/labstack/echo/v4"
)

func RegisterAdminController(c echo.Context) error {
	admin := models.Admin{}
	c.Bind(&admin)

	if err := configs.DB.Save(&admin).Error; err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseNotData{
			Status:  "Proses Register Gagal",
			Message: "Terdeteksi Eror",
		})
	}

	return c.JSON(http.StatusOK, helpers.ResponseData{
		Status:  "Proses Registrasi Berhasil",
		Message: "Successfuly Registrasi Admin",
	})

}

func LoginAdminController(c echo.Context) error {
	admin := models.Admin{}
	c.Bind(&admin)

	result, err := database.LoginAdmin(admin)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	reflectValue := reflect.ValueOf(result)
	adminID := reflectValue.FieldByName("ID").Interface().(uint)
	adminName := reflectValue.FieldByName("Name").Interface().(string)
	adminEmail := reflectValue.FieldByName("Email").Interface().(string)
	adminContact := reflectValue.FieldByName("Contact").Interface().(string)
	adminRole := reflectValue.FieldByName("Role").Interface().(string)

	token, err := middlewares.CreateTokenAdmin(int(adminID), adminName, adminEmail, adminContact, adminRole)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	adminResponse := models.AdminResponse{ID: adminID, Name: adminName, Email: adminEmail, Role: adminRole, Token: token}

	return c.JSON(http.StatusOK, helpers.ResponseData{
		Status:  "Proses Login Berhasil",
		Message: "Successfuly Login Admin",
		Data:    adminResponse,
	})
}

func GetAdminControllerAll(c echo.Context) error {
	var adminAll []models.Admin

	if err := configs.DB.Find(&adminAll).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, helpers.ResponseData{
		Status:  "Success Get Admin Data All",
		Message: "Successfuly",
		Data:    adminAll,
	})
}

func GetAdminById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Invalid Id",
		})
	}

	var admin models.Admin
	if err = configs.DB.Where("id = ?", id).First(&admin).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, helpers.ResponseData{
		Status:  "Success Get AdminProfile",
		Message: "Successfuly",
		Data:    admin,
	})
}

func UpdateAdminController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Invalid Id",
		})
	}

	var admin models.Admin
	if err := configs.DB.Where("id = ?", id).First(&admin).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]string{
			"message": "admin not found",
		})
	}

	c.Bind(&admin)
	if err := configs.DB.Model(&admin).Updates(admin).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, helpers.ResponseData{
		Status:  "Success Upload Admin Data",
		Message: "Successfuly",
		Data:    admin,
	})
}

func DeleteAdminController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Invalid Id",
		})
	}

	var admin models.Admin
	if err := configs.DB.First("id = ?", id).First(&admin).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]string{
			"message": "admin not found",
		})
	}

	if err := configs.DB.Delete(&admin).Error; err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseNotData{
			Status:  "Failed to Delete Admin Data",
			Message: "err.Error()",
		})
	}

	return c.JSON(http.StatusBadRequest, helpers.ResponseNotData{
		Status:  "Succes Delete Admin Data",
		Message: "Succesfuly",
	})
}

func GetUsersControllerAll(c echo.Context) error {
	var userAll []models.User

	if err := configs.DB.Find(&userAll).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, helpers.ResponseData{
		Status:  "Success Get Admin Data All",
		Message: "Successfuly",
		Data:    userAll,
	})
}

func GetUsersId(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Invalid Id",
		})
	}

	var user models.User
	if err = configs.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, helpers.ResponseData{
		Status:  "Success Get User by ID",
		Message: "Successfuly",
		Data:    user,
	})
}
