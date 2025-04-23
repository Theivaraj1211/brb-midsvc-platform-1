package router

import (
	"brb-midsvc-platform/controller"
	"brb-midsvc-platform/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(
	bookingCtrl *controller.BookingController,
	serviceCtrl *controller.ServicesController,
	vendorCtrl *controller.VendorController,
) *gin.Engine {
	r := gin.Default()

	// Admin routes
	admin := r.Group("/admin")
	admin.Use(middleware.AuthMiddleware("admin"))
	{
		// Vendor management
		admin.POST("/vendors", vendorCtrl.CreateVendor)
		admin.PUT("/vendors/:id", vendorCtrl.UpdateVendor)
		admin.PATCH("/vendors/:id/toggle", vendorCtrl.ToggleVendorStatus)

		// Service management
		admin.POST("/services", serviceCtrl.CreateService)
		admin.PUT("/services/:id", serviceCtrl.UpdateService)
		admin.PATCH("/services/:id/toggle", serviceCtrl.ToggleServiceStatus)
		admin.POST("/vendors/:vendorId/services/:serviceId/link", serviceCtrl.LinkToVendor)

		// Booking management
		admin.POST("/bookings", bookingCtrl.CreateBooking)
		admin.PUT("/bookings/:id/status", bookingCtrl.UpdateBookingStatus)
		admin.GET("/summary/vendor/:id", bookingCtrl.GetVendorSummary)
	}

	// Customer routes
	customer := r.Group("/customer")
	customer.Use(middleware.AuthMiddleware("customer"))
	{
		customer.POST("/bookings", bookingCtrl.CreateBooking)
		customer.GET("/bookings", bookingCtrl.GetCustomerBookings)
	}

	// Public routes (for demo purposes)
	public := r.Group("/api")
	{
		// public.GET("/vendors", vendorCtrl.ListVendors)
		// public.GET("/services", serviceCtrl.ListServices)
		public.GET("/summary/vendor/:id", bookingCtrl.GetVendorSummary)
	}

	return r
}
