package main

import (
	"assigment_7/database"
	"assigment_7/handler"
	"assigment_7/order"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("orders.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	database.Migrate(db)

	orderRepo := order.NewOrderRepository(db)
	orderSvc := order.NewOrderService(orderRepo)
	orderHandler := handler.NewOrderHandler(orderSvc)

	router := gin.Default()
	app := handler.NewRouter(router, orderHandler)
	app.Start(":8080")
}
