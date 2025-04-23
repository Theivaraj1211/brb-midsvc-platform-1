package services

import (
	"brb-midsvc-platform/model"
	"errors"
	"log"
	"time"

	"gorm.io/gorm"
)

type BookingService struct {
	db *gorm.DB
}

func NewBookingService(db *gorm.DB) *BookingService {
	return &BookingService{db: db}
}

func (s *BookingService) CreateBooking(booking *model.Booking) error {
	startTime, err := time.Parse(time.RFC3339, booking.StartTime)
	if err != nil {
		return errors.New("invalid start time format")
	}

	endTime, err := time.Parse(time.RFC3339, booking.EndTime)
	if err != nil {
		return errors.New("invalid end time format")
	}

	if !isValidBookingTime(startTime, endTime) {
		return errors.New("booking must be between 9AM-5PM in 1-hour slots")
	}

	var count int64
	s.db.Model(&model.Booking{}).
		Where("vendor_id = ?", booking.VendorID).
		Where("(start_time, end_time) OVERLAPS (?, ?)", booking.StartTime, booking.EndTime).
		Count(&count)

	if count > 0 {
		return errors.New("vendor has overlapping booking")
	}

	// Simulate notification
	go func() {
		log.Printf("Notification: New booking created for vendor %d at %s", booking.VendorID, booking.StartTime)
	}()

	return s.db.Create(booking).Error
}

func (s *BookingService) GetVendorSummary(vendorID string) (map[string]interface{}, error) {
	var totalBookings int64
	var statusCounts []struct {
		Status string
		Count  int64
	}

	if err := s.db.Model(&model.Booking{}).
		Where("vendor_id = ?", vendorID).
		Count(&totalBookings).Error; err != nil {
		return nil, err
	}

	if err := s.db.Model(&model.Booking{}).
		Select("status, count(*) as count").
		Where("vendor_id = ?", vendorID).
		Group("status").
		Scan(&statusCounts).Error; err != nil {
		return nil, err
	}

	statusMap := make(map[string]int64)
	for _, sc := range statusCounts {
		statusMap[sc.Status] = sc.Count
	}

	return map[string]interface{}{
		"vendor_id":          vendorID,
		"total_bookings":     totalBookings,
		"bookings_by_status": statusMap,
	}, nil
}

func (s *BookingService) UpdateBookingStatus(bookingID string, status model.BookingStatus) error {
	return s.db.Model(&model.Booking{}).
		Where("id = ?", bookingID).
		Update("status", status).Error
}

func isValidBookingTime(start, end time.Time) bool {
	return start.Hour() >= 9 && end.Hour() <= 17 &&
		end.Sub(start) == time.Hour
}
func (s *BookingService) GetCustomerBookings(customerID string) ([]model.Booking, error) {
	var bookings []model.Booking
	if err := s.db.Preload("Service").Preload("Vendor").
		Where("customer_id = ?", customerID).
		Find(&bookings).Error; err != nil {
		return nil, err
	}
	return bookings, nil
}
