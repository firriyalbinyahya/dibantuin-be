package controller

import (
	"dibantuin-be/entity"
	"dibantuin-be/service"
	"dibantuin-be/utils/response"
	"errors"
	"fmt"
	"math"
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

func (dpc *DonationProgramController) ListDonationPrograms(c *gin.Context) {
	statusRequest := c.Query("statusRequest")
	search := c.Query("search")

	limitStr := c.DefaultQuery("limit", "10")
	pageStr := c.DefaultQuery("page", "1")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 10
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		page = 1
	}

	categoryIDStr := c.Query("categoryId")
	var categoryID uint64
	if categoryIDStr != "" {
		parsedCategoryID, err := strconv.ParseUint(categoryIDStr, 10, 64)
		if err != nil {
			response.BuildErrorResponse(c, fmt.Errorf("invalid category_id"))
			return
		}
		categoryID = parsedCategoryID
	}

	programs, total, err := dpc.Service.ListDonationPrograms(statusRequest, search, limit, page, categoryID)
	if err != nil {
		response.BuildErrorResponse(c, err)
		return
	}

	response.BuildSuccessResponse(c, http.StatusOK, "Success to get list of programs", map[string]interface{}{
		"data":       programs,
		"total":      total,
		"page":       page,
		"limit":      limit,
		"total_page": int(math.Ceil(float64(total) / float64(limit))),
	}, nil)
}

func (dpc *DonationProgramController) GetDonationProgramDetail(c *gin.Context) {
	id := c.Param("id")
	programID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		response.BuildErrorResponse(c, fmt.Errorf("invalid program_id"))
		return
	}

	program, err := dpc.Service.GetDonationProgramDetail(programID)
	if err != nil {
		response.BuildErrorResponse(c, fmt.Errorf("failed to fetch data detail donation program"))
		return
	}

	response.BuildSuccessResponse(c, http.StatusOK, "Success get data detail donation program",
		program, nil)
}
