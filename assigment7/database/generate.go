package database

import (
	"time"

	"gorm.io/gorm"
)

type Orders struct {
	gorm.Model
	OrderID      int       `json:"orderId" gorm:"primary_key"`
	CustomerName string    `json:"customerName"`
	OrderedAt    time.Time `json:"orderedAt"`
	Items        []Items   `json:"items" gorm:"foreignkey:OrderID"`
}

type Items struct {
	gorm.Model
	ItemId      int    `json:"itemId" gorm:"primary_key"`
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    uint   `json:"quantity"`
	OrderID     uint   `json:"-"`
}

func Migratel(db *gorm.DB) {
	db.AutoMigrate(&Orders{}, &Items{})
}
