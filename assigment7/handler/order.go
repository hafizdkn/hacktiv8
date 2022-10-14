package handler

import (
	"github.com/gin-gonic/gin"

	"assigment_7/helper"
	"assigment_7/order"
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
		helper.WriteJsonRespnse(ctx, helper.BadResponse(err))
	}
	response := h.orderSvc.CreateOrder(input)
	helper.WriteJsonRespnse(ctx, response)
}

func (h *orderHandler) GetOrders(ctx *gin.Context) {
	response := h.orderSvc.GetOrders()
	helper.WriteJsonRespnse(ctx, response)
}

func (h *orderHandler) UpdateOrder(ctx *gin.Context) {
	var input order.UpdateOrder
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		helper.WriteJsonRespnse(ctx, helper.BadResponse(err))
	}
	response := h.orderSvc.UpdateOrder(input)
	helper.WriteJsonRespnse(ctx, response)
}
