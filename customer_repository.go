package salty

import (
	"context"
	"github.com/yakiza/salty/internal"
)

type CustomerRepository interface {
	// Create interface responsible for the creation of a customer
	Create(ctx context.Context, customer internal.Customer) error
}
