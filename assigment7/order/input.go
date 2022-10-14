package order

import "time"

type CreateOderInput struct {
	CustomerName string                   `json:"customerName"`
	Items        []map[string]interface{} `json:"items"`
	OrderedAt    time.Time                `json:"orderedAt"`
}
type UpdateOrder struct {
	OrderId      int                      `json:"orderId"`
	CustomerName string                   `json:"customerName"`
	OrderedAt    time.Time                `json:"orderedAt"`
	Items        []map[string]interface{} `json:"items"`
}
