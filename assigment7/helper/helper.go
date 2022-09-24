package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Payload interface{} `json:"payload,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

func WriteJsonRespnse(ctx *gin.Context, resp *Response) {
	ctx.JSON(resp.Status, resp)
}

func SuccessCreateResponse(payload interface{}, message string) *Response {
	return &Response{
		Status:  http.StatusCreated,
		Message: message,
		Payload: payload,
	}
}

func SuccessResponse(payload interface{}, message string) *Response {
	return &Response{
		Status:  http.StatusOK,
		Message: message,
		Payload: payload,
	}
}

func InternalServerError(err error) *Response {
	return &Response{
		Status:  http.StatusInternalServerError,
		Message: "INTERNAL_SERVER_ERROR",
		Error:   err.Error(),
	}
}
