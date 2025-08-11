package controller

import "dibantuin-be/service"

type DonationProgramController struct {
	Service *service.DonationProgramService
}

func NewDonationProgramController(service *service.DonationProgramService) *DonationProgramController {
	return &DonationProgramController{Service: service}
}
