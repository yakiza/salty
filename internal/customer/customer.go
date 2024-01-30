package customer

import (
	"context"
	"github.com/yakiza/salty/internal"
)

type CreateCustomerUseCase struct {
	customer internal.CustomerRepository
}

func (uc CreateCustomerUseCase) Create(ctx context.Context, customer internal.Customer) error {
	err := uc.customer.Create(ctx, customer)
	if err != nil {
		return err
	}

	return nil
}

func NewCreateCustomerUseCase(db internal.CustomerRepository) CreateCustomerUseCase {

	return CreateCustomerUseCase{}
}
