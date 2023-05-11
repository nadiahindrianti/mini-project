package models

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Contact  string `json:"contact"`
	Role     string `json:"role"`
	Password string `json:"password"`
}

type AdminResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
	Token string `json:"token"`
}

type AdminResponseStand struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Contact string `json:"contact"`
	Role    string `json:"role"`
}
