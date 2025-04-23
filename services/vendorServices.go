package services

import (
	"brb-midsvc-platform/model"
	"errors"
	"strconv"

	"gorm.io/gorm"
)

type VendorService struct {
	db *gorm.DB
}

func NewVendorService(db *gorm.DB) *VendorService {
	return &VendorService{db: db}
}

func (s *VendorService) CreateVendor(vendor *model.Vendor) error {
	return s.db.Create(vendor).Error
}

func (s *VendorService) UpdateVendor(vendorID string, vendor *model.Vendor) error {
	id, err := strconv.ParseUint(vendorID, 10, 32)
	if err != nil {
		return errors.New("invalid vendor ID")
	}
	vendor.ID = uint(id)
	return s.db.Save(vendor).Error
}

func (s *VendorService) ToggleVendorStatus(vendorID string) error {
	var vendor model.Vendor
	if err := s.db.First(&vendor, vendorID).Error; err != nil {
		return err
	}

	return s.db.Model(&vendor).
		Update("active", !vendor.Active).Error
}
