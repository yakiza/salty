package postgresql

import (
	"context"
	"github.com/yakiza/salty"
	"github.com/yakiza/salty/internal"
)

var _ salty.CustomerRepository = &CustomerRepository{}

type CustomerRepository struct {
	db *DB
}

func (c CustomerRepository) Create(ctx context.Context, customer internal.Customer) error {
	//TODO implement me
	panic("implement me")
}

func NewCustomerRepository(db *DB) CustomerRepository {
	return CustomerRepository{db}
}
