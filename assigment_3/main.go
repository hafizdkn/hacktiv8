package main

import (
	"github.com/gin-gonic/gin"

	"project_3/handler"
	"project_3/weather"
)

func main() {
	weaterSvc := weather.NewService()
	weatherHandler := handler.NewHandler(weaterSvc)

	router := gin.Default()
	app := handler.NewRouter(router, weatherHandler)
	app.Start(":8080")
}
