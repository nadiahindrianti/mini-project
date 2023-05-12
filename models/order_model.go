package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	ProductID   int     `json:"product_id" form:"product_id"`
	Product     Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserID      int     `json:"user_id" form:"user_id"`
	User        User    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ProductName string  `json:"product_name" form:"product_name"`
	JumlahOrder int     `json:"jumlah_order" form:"jumlah_order"`
}
