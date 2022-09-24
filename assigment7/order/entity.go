package order

import "time"

type Order struct {
	OrderID      int       `json:"orderId"`
	CustomerName string    `json:"customerName"`
	OrderedAt    time.Time `json:"orderedAt"`
}

type Item struct {
	ItemId      int    `json:"itemId,omitempty"`
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	OrderID     int    `json:"orderId"`
}

type Orders struct {
	Order Order
	Items Item
}
