package handler

import "github.com/gin-gonic/gin"

type routers struct {
	router  *gin.Engine
	weather *WeatherHandler
}

func NewRouter(router *gin.Engine, weatherHandler *WeatherHandler) *routers {
	return &routers{router, weatherHandler}
}

func (r *routers) Start(port string) {
	r.router.LoadHTMLGlob("templates/*.html")
	r.router.GET("/location", r.weather.GetLocation)

	r.router.Run(port)
}
