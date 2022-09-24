package order

import "time"

type Order struct {
	OrderID      int
	CustomerName string
	OrderedAt    time.Time
	Items        []Item
}

type Item struct {
	ItemId      int
	ItemCode    string
	Description string
	Quantity    int
	OrderID     int
}
