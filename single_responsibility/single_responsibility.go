package singleresponsibility

import "time"

// Single responsibility says "each class should be responsible for one thing" in we can do that in go by
// having different structs

// Creating the entity for order
type Order struct {
	OrderId    int
	OrderTotal float64
	CreatedAt  time.Time
}

// Defining the interface
type IOrderStore interface {
	CreateOrder() Order
}

// Defining the store and its dependencies
type OrderStore struct {
	// define the dependencies
}

func NewOrderStore() IOrderStore {
	return &OrderStore{}
}

func (*OrderStore) CreateOrder() Order {
	return Order{
		OrderId:    1,
		OrderTotal: 23.99,
		CreatedAt:  time.Now(),
	}
}
