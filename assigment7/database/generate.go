package database

import (
	"time"

	"gorm.io/gorm"
)

type Orders struct {
	gorm.Model
	OrderID      int
	CustomerName string
	OrderedAt    time.Time
}

type Items struct {
	gorm.Model
	LineItemId  int
	ItemId      int
	ItemCode    string
	Description string
	Quantity    int
	OrderID     int
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&Orders{}, &Items{})
}
