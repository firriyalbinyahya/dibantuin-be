package controller

import (
	"dibantuin-be/entity"
	"dibantuin-be/service"
	"dibantuin-be/utils/response"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DonationProgramController struct {
	Service *service.DonationProgramService
}

func NewDonationProgramController(service *service.DonationProgramService) *DonationProgramController {
	return &DonationProgramController{Service: service}
}

func (dpc *DonationProgramController) RequestProgram(c *gin.Context) {
	var donationProgramCreate entity.DonationProgramRequestCreate

	if err := c.ShouldBindJSON(&donationProgramCreate); err != nil {
		response.BuildErrorResponse(c, err)
		return
	}

	idUser, ok := c.Get("user_id")
	if !ok {
		response.BuildErrorResponse(c, errors.New("user id did not valid"))
		return
	}

	idUserInt, ok := idUser.(uint64)
	if !ok {
		response.BuildErrorResponse(c, errors.New("user id must uint64"))
		return
	}

	err := dpc.Service.CreateRequest(&donationProgramCreate, idUserInt)
	if err != nil {
		response.BuildErrorResponse(c, err)
		return
	}

	response.BuildSuccessResponse(c, http.StatusCreated, "Success request donation program. Wait admin to verify.", donationProgramCreate, nil)
}

func (dpc *DonationProgramController) VerifyProgram(c *gin.Context) {

	programRequestIDParam := c.Param("id")
	programRequestID, err := strconv.ParseUint(programRequestIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid program request id"})
		return
	}

	adminID := c.MustGet("user_id").(uint64)

	var verificationProgramRequest *entity.VerificationProgramRequest
	if err := c.ShouldBindJSON(&verificationProgramRequest); err != nil {
		response.BuildErrorResponse(c, err)
		return
	}

	err = dpc.Service.VerifyProgram(programRequestID, uint64(adminID), verificationProgramRequest.Status, verificationProgramRequest.Note)
	if err != nil {
		response.BuildErrorResponse(c, err)
		return
	}

	response.BuildSuccessResponse(c, http.StatusOK, "Verification updated", verificationProgramRequest, nil)

}
