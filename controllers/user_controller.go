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

func RegisterUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	if err := configs.DB.Save(&user).Error; err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseNotData{
			Status:  "Proses Register Gagal",
			Message: "Terdeteksi Eror",
		})
	}

	return c.JSON(http.StatusOK, helpers.ResponseNotData{
		Status:  "Proses Registrasi Berhasil",
		Message: "Successfuly Registrasi User",
	})

}

func LoginUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	result, err := database.LoginUser(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	reflectValue := reflect.ValueOf(result)
	userID := reflectValue.FieldByName("ID").Interface().(uint)
	userName := reflectValue.FieldByName("Name").Interface().(string)
	userEmail := reflectValue.FieldByName("Email").Interface().(string)
	userContact := reflectValue.FieldByName("Contact").Interface().(string)
	userAlamat := reflectValue.FieldByName("Alamat").Interface().(string)
	userRole := reflectValue.FieldByName("Role").Interface().(string)

	token, err := middlewares.CreateTokenUser(int(userID), userName, userEmail, userContact, userAlamat, userRole)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	userResponse := models.UserResponse{ID: userID, Name: userName, Email: userEmail, Role: userRole, Token: token}

	return c.JSON(http.StatusOK, helpers.ResponseData{
		Status:  "Proses Login Berhasil",
		Message: "Successfuly Login User",
		Data:    userResponse,
	})
}

func GetUserProfile(c echo.Context) error {
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

func UpdateUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseNotData{
			Status:  "Error",
			Message: "Data Tidak Tersedia",
		})
	}
	var user models.User
	if err := configs.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseNotData{
			Status:  "Error",
			Message: "Data Tidak Tersedia",
		})
	}

	c.Bind(&user)
	if err := configs.DB.Model(&user).Updates(user).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	userResp := models.UserResponseStand{
		Name:    user.Name,
		Email:   user.Email,
		Contact: user.Contact,
		Alamat:  user.Email,
		Role:    user.Role,
	}

	return c.JSON(http.StatusOK, helpers.ResponseData{
		Status:  "Success Upload User Data",
		Message: "Successfuly",
		Data:    userResp,
	})

}

func DeleteUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseNotData{
			Status:  "Error",
			Message: "Data Tidak Tersedia",
		})
	}

	var user models.User
	if err := configs.DB.First(&user, "id = ? ", id).Error; err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseNotData{
			Status:  "Error",
			Message: "Data Tidak Tersedia",
		})
	}

	if err := configs.DB.Delete(&user).Error; err != nil {
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
