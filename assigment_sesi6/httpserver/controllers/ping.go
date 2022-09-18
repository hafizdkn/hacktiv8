package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// type PingHandler struct {
// }

// func NewPingHandler() *PingHandler {
// 	return &PingHandler{}
// }

// func (p *PingHandler) HealthCheck(ctx *gin.Context) {
// 	ctx.JSON(http.StatusOK, map[string]interface{}{
// 		"health": "ok",
// 	})
// }

func HealthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"health": "ok",
	})
}
