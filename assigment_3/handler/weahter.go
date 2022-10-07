package handler

import (
	"github.com/gin-gonic/gin"

	"project_3/helper"
	"project_3/weather"
)

type WeatherHandler struct {
	weatherSvc weather.Service
}

func NewHandler(weatherSvc weather.Service) *WeatherHandler {
	return &WeatherHandler{weatherSvc}
}

func (h *WeatherHandler) GetLocation(ctx *gin.Context) {
	response := h.weatherSvc.GetLocation()
	helper.RenderHtmlResponse(ctx, "index.html", response.Payload)
}
