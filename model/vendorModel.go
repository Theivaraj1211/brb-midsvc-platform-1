package model

import "gorm.io/gorm"

type Vendor struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Description string
	Active      bool      `gorm:"default:true"`
	Services    []Service `gorm:"foreignKey:VendorID"`
}
