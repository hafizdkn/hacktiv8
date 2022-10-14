package handler

import (
	"github.com/gin-gonic/gin"
)

type routers struct {
	router       *gin.Engine
	orderHandler *orderHandler
}

func NewRouter(router *gin.Engine, orderHandler *orderHandler) *routers {
	return &routers{router, orderHandler}
}

func (r *routers) Start(port string) {
	r.router.POST("/order", r.orderHandler.CreateOrder)
	r.router.GET("/orders", r.orderHandler.GetOrders)
	r.router.POST("/update", r.orderHandler.UpdateOrder)

	r.router.Run(port)
}
