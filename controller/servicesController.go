package controller

import (
	"brb-midsvc-platform/model"
	"brb-midsvc-platform/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ServicesController struct {
	serviceService *services.ServicesService
}

func NewServicesController(ss *services.ServicesService) *ServicesController {
	return &ServicesController{serviceService: ss}
}

func (c *ServicesController) CreateService(ctx *gin.Context) {
	var service model.Service
	if err := ctx.ShouldBindJSON(&service); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.serviceService.CreateService(&service); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, service)
}

func (c *ServicesController) UpdateService(ctx *gin.Context) {
	serviceID := ctx.Param("id")
	var service model.Service
	if err := ctx.ShouldBindJSON(&service); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.serviceService.UpdateService(serviceID, &service); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, service)
}

func (c *ServicesController) ToggleServiceStatus(ctx *gin.Context) {
	serviceID := ctx.Param("id")
	if err := c.serviceService.ToggleServiceStatus(serviceID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Service status toggled"})
}

func (c *ServicesController) LinkToVendor(ctx *gin.Context) {
	vendorID := ctx.Param("vendorId")
	serviceID := ctx.Param("serviceId")

	if err := c.serviceService.LinkToVendor(serviceID, vendorID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Service linked to vendor"})
}
