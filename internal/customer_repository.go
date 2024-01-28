package internal

import "context"

type CustomerRepository interface {
	// Create interface responsible for the creation of a customer
	Create(ctx context.Context, customer Customer) error
}
