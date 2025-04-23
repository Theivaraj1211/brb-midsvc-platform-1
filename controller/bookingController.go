package controller

import (
	"brb-midsvc-platform/model"
	"brb-midsvc-platform/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookingController struct {
	bookingService *services.BookingService
}

func NewBookingController(bs *services.BookingService) *BookingController {
	return &BookingController{bookingService: bs}
}

func (c *BookingController) CreateBooking(ctx *gin.Context) {
	var booking model.Booking
	if err := ctx.ShouldBindJSON(&booking); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.bookingService.CreateBooking(&booking); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, booking)
}

func (c *BookingController) GetVendorSummary(ctx *gin.Context) {
	vendorID := ctx.Param("id")
	summary, err := c.bookingService.GetVendorSummary(vendorID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, summary)
}

func (c *BookingController) UpdateBookingStatus(ctx *gin.Context) {
	bookingID := ctx.Param("id")
	var statusUpdate struct {
		Status model.BookingStatus `json:"status"`
	}

	if err := ctx.ShouldBindJSON(&statusUpdate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.bookingService.UpdateBookingStatus(bookingID, statusUpdate.Status); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Booking status updated"})
}
func (c *BookingController) GetCustomerBookings(ctx *gin.Context) {
	// Get customer ID from JWT token
	customerID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "customer ID not found in token"})
		return
	}

	// Convert to string
	customerIDStr, ok := customerID.(string)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid customer ID format"})
		return
	}

	bookings, err := c.bookingService.GetCustomerBookings(customerIDStr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, bookings)
}
