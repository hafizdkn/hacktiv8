package main

import (
	"sesi6-gin/httpserver"
)

func main() {
	// router := gin.Default()

	// pingRepo := repo.NewPingRepo()
	// pingSvc := svc.NewPingSvc(pingRepo)
	// pingHandler := controllers.NewPingHandler(pingSvc)

	// userRepo := repo.NewUserRepo()
	// userSvc := svc.NewUserSvc(userRepo)
	// userHandler := controllers.NewUserHandler(userSvc)

	// app := httpserver.NewRouter(router, pingHandler, userHandler)
	// app.Start(":4000")

	app := httpserver.CreateRouter()
	app.Run(":4000")
}
