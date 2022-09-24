package order

import (
	"fmt"
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
	fmt.Printf("%+v\n", input)

	order.CustomerName = input.CustomerName
	order.OrderedAt = input.OrderedAt
	for _, item := range input.Items {
		order.Items = []Item{{
			Description: item["description"].(string),
			ItemCode:    item["itemCode"].(string),
			Quantity:    int(item["quantity"].(float64)),
		}}
	}
	fmt.Printf("%+v\n", order)
	return &order
}
