package controllers

import (
	"net/http"

	"sesi6-gin/httpserver/controllers/params"
	"sesi6-gin/httpserver/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CreateUser(ctx *gin.Context) {
	var req params.UserCreateRequest

	// step : (2) nge parsing data dari client
	// dan dimasukin ke params
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = validator.New().Struct(req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// step : (3) ngirim data ke services
	response := services.CreateUser(&req)

	// step : (9) kirim ke client
	WriteJsonRespnse(ctx, response)
}

func GetUser(ctx *gin.Context) {
	response := services.GetUser()
	WriteJsonRespnse(ctx, response)
}
