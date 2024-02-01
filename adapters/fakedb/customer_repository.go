package fakedb

import (
	"context"
	"github.com/yakiza/salty"
	"github.com/yakiza/salty/internal"
	"sync"
)

var _ salty.CustomerRepository = &CustomerRepository{}

type CustomerRepository struct {
	mux       sync.Mutex
	customers map[internal.CustomerName]internal.Customer
}

// Create creates a customer and stores it in a local dictionary
func (r *CustomerRepository) Create(_ context.Context, customer internal.Customer) error {
	r.mux.Lock()
	defer r.mux.Unlock()

	if _, ok := r.customers[customer.Name]; !ok {
		r.customers[customer.Name] = customer
	}

	return nil
}

func NewCustomerRepository() *CustomerRepository {
	return &CustomerRepository{
		mux:       sync.Mutex{},
		customers: make(map[internal.CustomerName]internal.Customer),
	}
}
