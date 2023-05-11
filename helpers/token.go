package helpers

import (
	"Mini_Project_Toko-Online/models"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

var mySigningKey = []byte("secret")

type MyCustomClaims struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.RegisteredClaims
}

func CreateToken(user *models.User) (string, error) {
	claims := MyCustomClaims{
		int(user.ID),
		user.Name,
		user.Email,
		user.Role,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(12 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)

	return ss, err
}

func ValidateToken(tokenString string) (any, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("token error")
	}

	claims, ok := token.Claims.(*MyCustomClaims)

	if !ok || !token.Valid {
		return nil, fmt.Errorf("token tidak ditemukan")
	}

	return claims, nil
}

func IsUser(next echo.HandlerFunc) echo.HandlerFunc {
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
			return echo.NewHTTPError(http.StatusUnauthorized, "user is not an admin")
		}
		return next(c)
	}
}
