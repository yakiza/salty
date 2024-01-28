package customer

import (
	"context"
	"github.com/yakiza/salty/internal"
)

type CreateCustomerUseCase struct {
	internal.CustomerRepository
}

func (uc CreateCustomerUseCase) Create(ctx context.Context, customer internal.Customer) error {

	return nil
}

func NewCreateCustomerUseCase() CreateCustomerUseCase {
	return CreateCustomerUseCase{}
}
