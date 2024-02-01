package customer

import (
	"context"
	"github.com/yakiza/salty"
	"github.com/yakiza/salty/internal"
	"go.uber.org/zap"
)

type UseCase struct {
	customer salty.CustomerRepository
	log      *zap.Logger
}

func (uc UseCase) Create(ctx context.Context, customer internal.Customer) error {
	err := uc.customer.Create(ctx, customer)
	// TODO: can wrap this on a logging package
	defer func() {
		if err == nil {
			uc.log.Info("Customer created successfully", zap.Any("customer", customer))
		} else {
			uc.log.Error("Create customer failure", zap.Error(err))
		}
	}()
	if err != nil {
		return err
	}

	return nil
}

func NewCustomerUseCase(db salty.CustomerRepository, logger *zap.Logger) UseCase {
	return UseCase{
		customer: db,
		log:      logger}
}
