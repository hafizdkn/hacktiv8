package order

import (
	"encoding/json"

	"assigment_7/helper"
)

type Service interface {
	CreateOrder(input CreateOderInput) *helper.Response
	GetOrders() *helper.Response
	UpdateOrder(input UpdateOrder) *helper.Response
}

type service struct {
	repository Repository
}

func NewOrderService(repository *repository) *service {
	return &service{repository}
}

func (s *service) CreateOrder(input CreateOderInput) *helper.Response {
	order, item := parseInputToModelCreate(input)
	newOrder, err := s.repository.Save(order)
	if err != nil {
		return helper.InternalServerError(err)
	}
	newItem, err := s.repository.Save(item)
	if err != nil {
		return helper.InternalServerError(err)
	}

	formatterOder := FormatOrder(newOrder.(Order), newItem.(Item))
	return helper.SuccessCreateResponse(formatterOder, "success create data order")
}

func (s *service) GetOrders() *helper.Response {
	order := s.repository.GetOrders()
	orders := FormatOrders(order)
	response := helper.SuccessResponse(orders, "success get all data orders")
	return response
}

func (s *service) UpdateOrder(input UpdateOrder) *helper.Response {
	order, item := parseInputToModelUpdate(input)
	dataOrder, err := s.repository.UpdateOrder(order)
	if err != nil {
		return helper.InternalServerError(err)
	}
	dataItem, err := s.repository.UpdateOrder(item)
	if err != nil {
		return helper.InternalServerError(err)
	}
	formatterOder := FormatOrder(dataOrder.(Order), dataItem.(Item))
	return helper.SuccessResponse(formatterOder, "SUCCESS UPDATE DATA ORDER")
}

func parseInputToModelCreate(input CreateOderInput) (Order, Item) {
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
	return order, i
}

func parseInputToModelUpdate(input UpdateOrder) (Order, Item) {
	var order Order
	var items []Item

	order.OrderID = input.OrderId
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
		i.LineItemId = item.LineItemId
		i.ItemCode = item.ItemCode
		i.Description = item.Description
		i.Quantity = item.Quantity
	}
	return order, i
}
