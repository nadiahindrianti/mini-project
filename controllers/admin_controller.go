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

	return c.JSON(http.StatusOK, helpers.ResponseNotData{
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
		Message: "Successfuly Login User",
		Data:    adminResponse,
	})
}

func GetAdminControllerAll(c echo.Context) error {
	var users []models.UserResponseStand

	if err := configs.DB.Find(&users).Error; err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseNotData{
			Status:  "Error",
			Message: "Data Tidak Tersedia",
		})
	}

	return c.JSON(http.StatusOK, helpers.ResponseData{
		Status:  "Success Get AdminProfile",
		Message: "Successfuly",
		Data:    users,
	})
}

func GetAdminProfile(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseNotData{
			Status:  "Error",
			Message: "Data Tidak Tersedia",
		})
	}

	var admin models.Admin
	if err = configs.DB.Where("id = ?", id).First(&admin).Error; err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseNotData{
			Status:  "Error",
			Message: "Data Tidak Tersedia",
		})
	}

	adminResp := models.AdminResponseStand{
		Name:    admin.Name,
		Email:   admin.Email,
		Contact: admin.Contact,
		Role:    admin.Role,
	}

	return c.JSON(http.StatusOK, helpers.ResponseData{
		Status:  "Success Get AdminProfile",
		Message: "Successfuly",
		Data:    adminResp,
	})
}

func UpdateAdminController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseNotData{
			Status:  "Error",
			Message: "Data Tidak Tersedia",
		})
	}
	var admin models.Admin
	if err := configs.DB.Where("id = ?", id).First(&admin).Error; err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseNotData{
			Status:  "Error",
			Message: "Data Tidak Tersedia",
		})
	}

	c.Bind(&admin)
	if err := configs.DB.Model(&admin).Updates(admin).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	adminResp := models.AdminResponseStand{
		Name:    admin.Name,
		Email:   admin.Email,
		Contact: admin.Contact,
		Role:    admin.Role,
	}

	return c.JSON(http.StatusOK, helpers.ResponseData{
		Status:  "Success Upload Admin Data",
		Message: "Successfuly",
		Data:    adminResp,
	})

}

func DeleteAdminController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseNotData{
			Status:  "Error",
			Message: "Data Tidak Tersedia",
		})
	}

	var admin models.Admin
	if err := configs.DB.First(&admin, "id = ? ", id).Error; err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseNotData{
			Status:  "Error",
			Message: "Data Tidak Tersedia",
		})
	}

	if err := configs.DB.Delete(&admin).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Gagal Delete User Data",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusBadRequest, helpers.ResponseNotData{
		Status:  "Succes Delete User Data",
		Message: "Succesfuly",
	})
}

func GetUsersControllerAll(c echo.Context) error {
	var users []models.UserResponseStand

	if err := configs.DB.Find(&users).Error; err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseNotData{
			Status:  "Error",
			Message: "Data Tidak Tersedia",
		})
	}

	return c.JSON(http.StatusOK, helpers.ResponseData{
		Status:  "Success Get AdminProfile",
		Message: "Successfuly",
		Data:    users,
	})
}

func GetUsersId(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseNotData{
			Status:  "Error",
			Message: "Data Tidak Tersedia",
		})
	}

	var user models.User
	if err = configs.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseNotData{
			Status:  "Error",
			Message: "Data Tidak Tersedia",
		})
	}

	userResp := models.UserResponseStand{
		Name:    user.Name,
		Email:   user.Email,
		Contact: user.Contact,
		Alamat:  user.Email,
		Role:    user.Role,
	}

	return c.JSON(http.StatusOK, helpers.ResponseData{
		Status:  "Success Get UserProfile",
		Message: "Successfuly",
		Data:    userResp,
	})
}
