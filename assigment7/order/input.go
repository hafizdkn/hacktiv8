package order

import "time"

type CreateOderInput struct {
	CustomerName string                   `json:"customerName"`
	Items        []map[string]interface{} `json:"items"`
	OrderedAt    time.Time                `json:"orderedAt"`
}
