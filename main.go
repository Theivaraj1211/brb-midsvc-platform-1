package main

import (
	"brb-midsvc-platform/config"
	"brb-midsvc-platform/controller"
	"brb-midsvc-platform/dbconnection"
	"brb-midsvc-platform/model"
	"brb-midsvc-platform/router"
	"brb-midsvc-platform/services"
	"log"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize database
	dbconnection.ConnectDB(cfg)
	db := dbconnection.DB

	// Auto-migrate models
	if err := db.AutoMigrate(
		&model.Vendor{},
		&model.Service{},
		&model.Booking{},
	); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Initialize services
	bookingService := services.NewBookingService(db)
	serviceService := services.NewServicesService(db)
	vendorService := services.NewVendorService(db)

	// Initialize controllers
	bookingCtrl := controller.NewBookingController(bookingService)
	servicesCtrl := controller.NewServicesController(serviceService)
	vendorCtrl := controller.NewVendorController(vendorService)

	// Setup router
	r := router.SetupRouter(bookingCtrl, servicesCtrl, vendorCtrl)

	// Start server
	log.Printf("Server running on port %s", cfg.DBPort)
	if err := r.Run(":" + cfg.DBPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
