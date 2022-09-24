package order

import (
	"gorm.io/gorm"
)

type Repository interface {
	Save(order Order, item Item) (Order, Item, error)
	GetOrders() []Orders
	// UpdateOrder(order Order) (Order, error)
}
type repository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(order Order, item Item) (Order, Item, error) {
	newOrder, newItem := r.generateId(order, item)
	err := r.db.Create(&newOrder).Error
	if err != nil {
		return Order{}, Item{}, err
	}

	err = r.db.Create(&newItem).Error
	if err != nil {
		return Order{}, Item{}, err
	}

	return newOrder, newItem, nil
}

func (r *repository) GetOrders() []Orders {
	var o []Order
	var i []Item

	r.db.Find(&o)
	r.db.Find(&i)

	var orders []Orders
	for index, order := range o {
		item := i[index]
		orders = append(orders, Orders{Order: order, Items: item})
	}
	return orders
}

func (r *repository) generateId(order Order, item Item) (Order, Item) {
	orders := make([]Order, 0)
	r.db.Find(&orders)

	id := 1
	lengthData := len(orders)
	if lengthData > 0 {
		id = lengthData + 1
	}
	order.OrderID = id
	item.OrderID = id

	return order, item
}
