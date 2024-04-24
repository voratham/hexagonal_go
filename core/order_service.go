package core

import "errors"

// primary port
type OrderService interface {
	CreateOrder(order Order) error
}

type orderServiceImpl struct {
	repo OrderRepository
}

func NewOrderService(repo OrderRepository) OrderService {
	return &orderServiceImpl{repo: repo}
}

// CreateOrder implements OrderService.
func (o *orderServiceImpl) CreateOrder(order Order) error {
	// Business logic function
	if order.Total <= 0 {
		return errors.New("total must be positive")
	}

	if err := o.repo.Save(order); err != nil {
		return err
	}

	return nil

}
