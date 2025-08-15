package controller

import "dibantuin-be/service"

type DonationReportController struct {
	DonationReportService *service.DonationReportService
}

func NewDonationReportController(service *service.DonationReportService) *DonationReportController {
	return &DonationReportController{DonationReportService: service}
}
