package model

import "time"

type StoreType string

const (
	Silver   StoreType = "silver"
	Gold     StoreType = "gold"
	Platinum StoreType = "platinum"
)

type Store struct {
	StoreID   uint      `json:"store_id" gorm:"primaryKey;autoIncrement"`
	Email     string    `json:"email"    gorm:"not null;unique"`
	Password  password  `json:"-"        gorm:"embedded;embeddedPrefix:password_"`
	Name      string    `json:"name"     gorm:"not null"`
	Type      StoreType `json:"type"     gorm:"not null"`
	CreatedAt time.Time `json:"-"        gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"-"        gorm:"autoUpdateTime"`

	// foreign key
	Products []Product `json:"products,omitempty" gorm:"foreignKey:StoreID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
