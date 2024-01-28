package fakedb

import (
	"context"
	"github.com/yakiza/salty/internal"
	"sync"
)

type CustomerRepository interface {
	Create(ctx context.Context) error
}

type CustomerRepositoryFakeDB struct {
	mux       sync.Mutex
	customers map[internal.CustomerName]internal.Customer
}

func (r *CustomerRepositoryFakeDB) Create(_ context.Context, customer internal.Customer) error {
	r.mux.Lock()
	defer r.mux.Unlock()

	if _, ok := r.customers[customer.Name]; !ok {
		r.customers[customer.Name] = customer
	}

	return nil
}
