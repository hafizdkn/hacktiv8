package handler

import (
	"fmt"

	"assigment_7/order"

	"github.com/gin-gonic/gin"
)

type orderHandler struct {
	orderSvc order.Service
}

func NewOrderHandler(orderSvc order.Service) *orderHandler {
	return &orderHandler{orderSvc}
}

func (h *orderHandler) CreateOrder(ctx *gin.Context) {
	var input order.CreateOderInput
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		panic(err)
	}
	response, err := h.orderSvc.CreateOrder(input)
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}
