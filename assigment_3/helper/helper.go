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

func RenderHtmlResponse(ctx *gin.Context, fileHtml string, content interface{}) {
	ctx.HTML(http.StatusOK, fileHtml, gin.H{
		"content": content,
	})
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

func ErrorResponse(err error) *Response {
	return &Response{
		Status:  http.StatusBadRequest,
		Message: "BAD REQUESTS",
		Error:   err.Error(),
	}
}
