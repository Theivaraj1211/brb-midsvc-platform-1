package services

import (
	"brb-midsvc-platform/model"
	"errors"
	"strconv"

	"gorm.io/gorm"
)

type ServicesService struct {
	db *gorm.DB
}

func NewServicesService(db *gorm.DB) *ServicesService {
	return &ServicesService{db: db}
}

func (s *ServicesService) CreateService(service *model.Service) error {
	return s.db.Create(service).Error
}

func (s *ServicesService) UpdateService(serviceID string, service *model.Service) error {
	id, err := strconv.ParseUint(serviceID, 10, 32)
	if err != nil {
		return errors.New("invalid service ID")
	}
	service.ID = uint(id)
	return s.db.Save(service).Error
}

func (s *ServicesService) ToggleServiceStatus(serviceID string) error {
	var service model.Service
	if err := s.db.First(&service, serviceID).Error; err != nil {
		return err
	}

	return s.db.Model(&service).
		Update("active", !service.Active).Error
}

func (s *ServicesService) LinkToVendor(serviceID, vendorID string) error {
	serviceUint, err := strconv.ParseUint(serviceID, 10, 32)
	if err != nil {
		return errors.New("invalid service ID")
	}

	vendorUint, err := strconv.ParseUint(vendorID, 10, 32)
	if err != nil {
		return errors.New("invalid vendor ID")
	}

	return s.db.Model(&model.Service{}).
		Where("id = ?", serviceUint).
		Update("vendor_id", vendorUint).Error
}
