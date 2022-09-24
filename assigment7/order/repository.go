package order

import "gorm.io/gorm"

type Repository interface {
	Save(order Order) (Order, error)
	// UpdateOrder(order Order) (Order, error)
	// GetOrder() ([]Order, error)
}
type repository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(order Order) (Order, error) {
	err := r.db.Create(&order).Error
	if err != nil {
		return Order{}, nil
	}
	return order, nil
}
