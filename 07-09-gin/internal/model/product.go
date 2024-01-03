package model

import "time"

type Product struct {
	ProductID   uint      `json:"product_id" gorm:"primaryKey;autoIncrement"`
	StoreID     uint      `json:"store_id"   gorm:"not null;foreignKey"`
	Name        string    `json:"name"       gorm:"not null"`
	Description string    `gorm:"not null"`
	ImageUrl    string    `gorm:"not null"`
	Price       int       `gorm:"not null;check:price > 0"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
