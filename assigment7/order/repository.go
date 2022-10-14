package order

import (
	"gorm.io/gorm"
)

type Repository interface {
	Save(order interface{}) (interface{}, error)
	GetOrders() []Orders
	UpdateOrder(order interface{}) (interface{}, error)
}
type repository struct {
	db       *gorm.DB
	OrderId  int
	UpdateId int
}

func NewOrderRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) Save(order interface{}) (interface{}, error) {
	var err error
	var dataOrder interface{}

	switch order.(type) {
	case Order:
		r.generateId(order)
		o := order.(Order)
		o.OrderID = r.OrderId
		dataOrder = o
		err = r.db.Create(&o).Error
		if err != nil {
			return nil, err
		}
	case Item:
		i := order.(Item)
		i.OrderID = r.OrderId
		dataOrder = i
		err = r.db.Create(&i).Error
		if err != nil {
			return nil, err
		}
	}
	return dataOrder, nil
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

func (r *repository) UpdateOrder(order interface{}) (interface{}, error) {
	var dataOrder interface{}
	switch order.(type) {
	case Order:
		o := order.(Order)
		err := r.db.Debug().Model(&Order{}).Where("order_id", o.OrderID).Updates(o).Error
		if err != nil {
			return nil, err
		}
		dataOrder = o
		r.UpdateId = o.OrderID
	case Item:
		i := order.(Item)
		err := r.db.Debug().Model(&Item{}).Where("order_id", r.UpdateId).Updates(i).Error
		if err != nil {
			return nil, err
		}
		dataOrder = i
	}
	return dataOrder, nil
}

func (r *repository) generateId(order interface{}) {
	var orderId int
	switch order.(type) {
	case Order:
		orders := make([]Order, 0)
		r.db.Find(&orders)
		id := 1
		lengthData := len(orders)
		if lengthData > 0 {
			id = lengthData + 1
		}
		orderId = id
	}
	r.OrderId = orderId
}
