package handler

import (
	"fmt"

	"assigment_7/helper"
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
	fmt.Println(input)
	if err != nil {
		panic(err)
	}
	response := h.orderSvc.CreateOrder(input)
	helper.WriteJsonRespnse(ctx, response)
}

func (h *orderHandler) GetOrders(ctx *gin.Context) {
	response := h.orderSvc.GetOrders()
	helper.WriteJsonRespnse(ctx, response)
}
