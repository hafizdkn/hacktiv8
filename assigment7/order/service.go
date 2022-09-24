package order

import (
	"fmt"
	"time"
)

type Service interface {
	CreateOrder(input CreateOderInput) (Order, error)
}

type service struct {
	repository Repository
}

func NewOrderService(repository *repository) *service {
	return &service{repository}
}

func (s *service) CreateOrder(input CreateOderInput) (Order, error) {
	newOrder := parserInputToModel(input)
	order, err := s.repository.Save(*newOrder)
	if err != nil {
		return Order{}, nil
	}
	return order, nil
}

func parserInputToModel(input CreateOderInput) *Order {
	var order Order
	for _, item := range input.Items {
		order.CustomerName = item["customerName"].(string)
		order.Items = item["items"].([]Item)
		order.OrderedAt = item["orderedAt"].(time.Time)

	}
	fmt.Printf("%+v\n", order)
	return &order
}
