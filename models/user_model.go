package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Contact  string `json:"contact"`
	Alamat   string `json:"alamat"`
	Role     string `json:"role"`
	Password string `json:"password"`
}

type UserResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
	Token string `json:"token"`
}

type UserResponseStand struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Contact string `json:"contact"`
	Alamat  string `json:"alamat"`
	Role    string `json:"role"`
}
