package middlewares

import (
	"Mini_Project_Toko-Online/constants"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func CreateTokenUser(ID int, name string, email string, contact string, alamat string, role string) (string, error) {
	claims := jwt.MapClaims{}
	claims["userID"] = ID
	claims["name"] = name
	claims["email"] = email
	claims["contact"] = contact
	claims["alamat"] = alamat
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 6).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.SECRET_JWT))
}

func CreateTokenAdmin(ID int, name string, email string, contact string, role string) (string, error) {
	claims := jwt.MapClaims{}
	claims["ID"] = ID
	claims["name"] = name
	claims["email"] = name
	claims["contact"] = name
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.SECRET_JWT))
}

func ExtractTokenUserId(e echo.Context) int {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["userID"].(int)
		return userId
	}
	return 0
}

func ExtractTokenAdminId(e echo.Context) int {
	user := e.Get("admin").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["adminID"].(int)
		return userId
	}
	return 0
}

func IsUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		admin, ok := c.Get("admin").(*jwt.Token)
		if !ok {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid or missing jwt token")
		}
		claims, ok := admin.Claims.(jwt.MapClaims)
		if !ok {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid jwt claims")
		}
		if role, ok := claims["role"].(string); !ok || role != "User" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Not Token for Admin")
		}
		return next(c)
	}
}

func IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user, ok := c.Get("user").(*jwt.Token)
		if !ok {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid or missing jwt token")
		}
		claims, ok := user.Claims.(jwt.MapClaims)
		if !ok {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid jwt claims")
		}
		if role, ok := claims["role"].(string); !ok || role != "Admin" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Not Token for User")
		}
		return next(c)
	}
}
