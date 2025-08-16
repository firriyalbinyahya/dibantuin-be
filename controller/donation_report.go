package controller

import (
	"dibantuin-be/service"

	"github.com/gin-gonic/gin"
)

type DonationReportController struct {
	DonationReportService *service.DonationReportService
}

func NewDonationReportController(service *service.DonationReportService) *DonationReportController {
	return &DonationReportController{DonationReportService: service}
}

func (drc *DonationReportController) CreateDonationReport(c *gin.Context) {
	
}
