package singleresponsibility

import "time"

// SRP states that a module should be responsible for one, and only one, reason to change.
// In Go, we can achieve this by creating focused structs and interfaces.

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
