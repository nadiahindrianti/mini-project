package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Title      string   `json:"title"`
	Price      float64  `json:"price"`
	Stock      float64  `json:"stock"`
	CategoryID int      `json:"category_id"`
	Category   Category `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
