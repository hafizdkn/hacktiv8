package httpserver

import (
	"sesi6-gin/httpserver/controllers"

	"github.com/gin-gonic/gin"
)

// type Router struct {
// 	router *gin.Engine
// 	ping   *controllers.PingHandler
// }

// func NewRouter(router *gin.Engine, ping *controllers.PingHandler) *Router {
// 	return &Router{
// 		router: router,
// 		ping:   ping,
// 	}
// }

// func (r *Router) Start(port string) {
// 	r.router.GET("/ping", r.ping.HealthCheck)

// 	r.router.Run(port)
// }

func CreateRouter() *gin.Engine {
	router := gin.Default()

	// step :(1) request masuk, request keluar
	router.GET("/ping", controllers.HealthCheck)
	router.POST("/users", controllers.CreateUser)
	router.GET("/users", controllers.GetUser)
	return router
}
