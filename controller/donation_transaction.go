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

type DonationTransactionController struct {
	Service *service.DonationTransactionService
}

func NewDonationTransactionController(service *service.DonationTransactionService) *DonationTransactionController {
	return &DonationTransactionController{Service: service}
}

func (dtc *DonationTransactionController) CreateMoneyDonationTransaction(c *gin.Context) {
	idUser, _ := c.Get("user_id")
	idUserInt, _ := idUser.(uint64)

	var request entity.MoneyTransactionDonationRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		response.BuildErrorResponse(c, err)
		return
	}

	transaction, err := dtc.Service.CreateMoneyDonationTransaction(&request, idUserInt)
	if err != nil {
		response.BuildErrorResponse(c, err)
		return
	}

	response.BuildSuccessResponse(c, http.StatusOK, "Success make donation. Wait admin to verify.", transaction, nil)
}

func (dtc *DonationTransactionController) VerifyTransactionDonation(c *gin.Context) {

	transactionDonationIDParam := c.Param("id")
	transactionDonationID, err := strconv.ParseUint(transactionDonationIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid program request id"})
		return
	}

	adminID := c.MustGet("user_id").(uint64)

	var verificationTransactionRequest *entity.VerificationTransactionRequest
	if err := c.ShouldBindJSON(&verificationTransactionRequest); err != nil {
		response.BuildErrorResponse(c, err)
		return
	}

	err = dtc.Service.VerifyDonationTransaction(transactionDonationID, uint64(adminID), verificationTransactionRequest.Status, verificationTransactionRequest.Note)
	if err != nil {
		response.BuildErrorResponse(c, err)
		return
	}

	response.BuildSuccessResponse(c, http.StatusOK, "Verification updated", verificationTransactionRequest, nil)
}

func (dtc *DonationTransactionController) ListDonationTransactions(c *gin.Context) {
	search := c.Query("search")
	status := c.Query("status")

	var filterUserID *uint64
	if userIDStr := c.Query("user_id"); userIDStr != "" {
		parsedID, err := strconv.ParseUint(userIDStr, 10, 64)
		if err != nil {
			response.BuildErrorResponse(c, errors.New("invalid user_id"))
			return
		}
		filterUserID = &parsedID
	}

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

	paginationData, err := dtc.Service.ListDonationTransactions(filterUserID, status, search, limit, page)
	if err != nil {
		response.BuildErrorResponse(c, err)
		return
	}

	response.BuildSuccessResponse(c, http.StatusOK, "Success get list of donation transactions",
		paginationData, nil)
}

func (dtc *DonationTransactionController) ListDonationTransactionsUser(c *gin.Context) {
	search := c.Query("search")
	status := c.Query("status")

	userID := c.MustGet("user_id").(uint64)

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

	paginationData, err := dtc.Service.ListDonationTransactions(&userID, status, search, limit, page)
	if err != nil {
		response.BuildErrorResponse(c, err)
		return
	}

	response.BuildSuccessResponse(c, http.StatusOK, "Success get list of donation transactions",
		paginationData, nil)
}
