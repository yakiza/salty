package salty

import (
	"context"
	"github.com/yakiza/salty/internal"
)

type CustomerUseCases interface {
	// Create use case executing the relevant domain actions to create a customer
	Create(ctx context.Context, customer internal.Customer) error
}
