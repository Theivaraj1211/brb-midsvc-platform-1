package controller

import (
	"brb-midsvc-platform/model"
	"brb-midsvc-platform/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type VendorController struct {
	vendorService *services.VendorService
}

func NewVendorController(vs *services.VendorService) *VendorController {
	return &VendorController{vendorService: vs}
}

func (c *VendorController) CreateVendor(ctx *gin.Context) {
	var vendor model.Vendor
	if err := ctx.ShouldBindJSON(&vendor); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.vendorService.CreateVendor(&vendor); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, vendor)
}

func (c *VendorController) UpdateVendor(ctx *gin.Context) {
	vendorID := ctx.Param("id")
	var vendor model.Vendor
	if err := ctx.ShouldBindJSON(&vendor); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.vendorService.UpdateVendor(vendorID, &vendor); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, vendor)
}

func (c *VendorController) ToggleVendorStatus(ctx *gin.Context) {
	vendorID := ctx.Param("id")
	if err := c.vendorService.ToggleVendorStatus(vendorID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Vendor status toggled"})
}
