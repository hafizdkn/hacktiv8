package order

import (
	"time"
)

type OrdersFormatter struct {
	OrderID      int       `json:"orderId"`
	CustomerName string    `json:"customerName"`
	OrderedAt    time.Time `json:"orderedAt"`
	Items        []Item    `json:"items"`
}

func FormatOrders(orders []Orders) *[]OrdersFormatter {
	var dataOrders []OrdersFormatter
	for _, order := range orders {
		item := order.Items
		dataOrders = append(dataOrders, OrdersFormatter{
			OrderID:      order.Order.OrderID,
			CustomerName: order.Order.CustomerName,
			OrderedAt:    order.Order.OrderedAt,
			Items:        []Item{item},
		})
	}
	return &dataOrders
}

func FormatOrder(order Order, item Item) *OrdersFormatter {
	return &OrdersFormatter{OrderID: order.OrderID, CustomerName: order.CustomerName, OrderedAt: order.OrderedAt, Items: []Item{item}}
}
