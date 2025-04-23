package model

import "gorm.io/gorm"

type BookingStatus string

const (
	Pending   BookingStatus = "pending"
	Confirmed BookingStatus = "confirmed"
	Completed BookingStatus = "completed"
)

type Booking struct {
	gorm.Model
	CustomerID uint
	ServiceID  uint
	Service    Service `gorm:"foreignKey:ServiceID"`
	VendorID   uint
	Vendor     Vendor        `gorm:"foreignKey:VendorID"`
	StartTime  string        `gorm:"not null"`
	EndTime    string        `gorm:"not null"`
	Status     BookingStatus `gorm:"type:varchar(20);default:'pending'"`
}
