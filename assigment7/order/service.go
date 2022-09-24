package order

import (
	"encoding/json"

	"assigment_7/helper"
)

type Service interface {
	CreateOrder(input CreateOderInput) *helper.Response
	GetOrders() *helper.Response
}

type service struct {
	repository Repository
}

func NewOrderService(repository *repository) *service {
	return &service{repository}
}

func (s *service) CreateOrder(input CreateOderInput) *helper.Response {
	newOrder, newItem := parseInputToModel(input)
	o, i, err := s.repository.Save(*newOrder, *newItem)
	order := FormatOrder(o, i)
	if err != nil {
		return helper.InternalServerError(err)
	}
	return helper.SuccessCreateResponse(order, "success create data order")
}

func (s *service) GetOrders() *helper.Response {
	order := s.repository.GetOrders()
	orders := FormatOrders(order)
	response := helper.SuccessResponse(orders, "success get all data orders")
	return response
}

func parseInputToModel(input CreateOderInput) (*Order, *Item) {
	var order Order
	var items []Item

	order.CustomerName = input.CustomerName
	order.OrderedAt = input.OrderedAt

	byteItem, err := json.Marshal(input.Items)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(byteItem, &items)
	if err != nil {
		panic(err)
	}
	var i Item
	for _, item := range items {
		i.ItemCode = item.ItemCode
		i.Description = item.Description
		i.Quantity = item.Quantity
	}
	return &order, &i
	// fmt.Printf("item %+v\n", item)
	// for _, item := range input.Items {
	// 	order.Items = []Item{{
	// 		Description: item["description"].(string),
	// 		ItemCode:    item["itemCode"].(string),
	// 		Quantity:    int(item["quantity"].(float64)),
	// 	}}
	// }
	// fmt.Printf("%+v\n", order)
}
