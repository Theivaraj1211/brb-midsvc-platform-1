package model

import "gorm.io/gorm"

type Service struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Description string
	Price       float64 `gorm:"not null"`
	Active      bool    `gorm:"default:true"`
	VendorID    uint    `gorm:"not null"`
	Vendor      Vendor  `gorm:"foreignKey:VendorID"`
}
