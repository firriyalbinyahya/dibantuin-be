package controller

import (
	"dibantuin-be/service"
	"dibantuin-be/utils/response"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type ReportController struct {
	ReportService *service.ReportService
}

func NewReportController(reportService *service.ReportService) *ReportController {
	return &ReportController{ReportService: reportService}
}

func (rc *ReportController) GetDonationProgramReport(c *gin.Context) {
	programIDParam := c.Param("id")
	programID, err := strconv.ParseUint(programIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid program id"})
		return
	}

	idUser, _ := c.Get("user_id")
	idUserInt, _ := idUser.(uint64)

	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	var start, end time.Time

	if startDate != "" {
		start, err = time.Parse("2006-01-02", startDate)
		if err != nil {
			response.BuildErrorResponse(c, errors.New("invalid start_date format, use YYYY-MM-DD"))
			return
		}
	}

	if endDate != "" {
		end, err = time.Parse("2006-01-02", endDate)
		if err != nil {
			response.BuildErrorResponse(c, errors.New("invalid start_date format, use YYYY-MM-DD"))
			return
		}
	}

	program, err := rc.ReportService.GetDonationProgramReport(programID, idUserInt, start, end)
	if err != nil {
		response.BuildErrorResponse(c, err)
		return
	}

	response.BuildSuccessResponse(c, http.StatusOK, "success to get data of donation program report", program, nil)
}

func (rc *ReportController) GetGlobalReport(c *gin.Context) {
	globalReport, err := rc.ReportService.GetGlobalReport()
	if err != nil {
		response.BuildErrorResponse(c, err)
		return
	}
	response.BuildSuccessResponse(c, http.StatusOK, "success to get data of global report", globalReport, nil)
}
