package response

import (
	"dibantuin-be/constants"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type JSONResponse struct {
	Code     string      `json:"code"`
	Message  string      `json:"message"`
	Data     interface{} `json:"data"`
	Metadata interface{} `json:"metadata,omitempty"`
}

func BuildSuccessResponse(c *gin.Context, statusCode int, message string, data any, metadata any) {
	c.JSON(statusCode, JSONResponse{
		Code:     constants.RESPONSE_SUCCESS,
		Message:  message,
		Data:     data,
		Metadata: metadata,
	})
}

func BuildErrorResponse(c *gin.Context, err error) {
	var customErr *CustomError
	var statusCode int
	var message string
	var errorDetails interface{}

	if errors.As(err, &customErr) {
		statusCode = customErr.StatusCode
		message = customErr.Message
		errorDetails = customErr.Err
	} else {
		statusCode = http.StatusInternalServerError
		message = "An unexpected error occurred"
		errorDetails = err.Error()
	}

	// map the status code to a response code
	var code string
	switch statusCode {
	case http.StatusBadRequest:
		code = constants.RESPONSE_BAD_REQUEST
	case http.StatusNotFound:
		code = constants.RESPONSE_NOT_FOUND
	case http.StatusUnauthorized:
		code = constants.RESPONSE_UNAUTHENTICATED
	case http.StatusForbidden:
		code = constants.RESPONSE_UNAUTHORIZED
	case http.StatusTooManyRequests:
		code = constants.RESPONSE_TOO_MANY_REQUESTS
	case http.StatusInternalServerError:
		code = constants.RESPONSE_INTERNAL_SERVER_ERROR
	default:
		code = constants.RESPONSE_INTERNAL_SERVER_ERROR
	}

	response := JSONResponse{
		Code:    code,
		Message: message,
		Data:    errorDetails,
	}

	c.JSON(statusCode, response)
}
