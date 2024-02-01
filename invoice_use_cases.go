package salty

import (
	"context"
)

type InvoiceUseCases interface {
	// New Is responsible in creating a new invoice entry and calling the appropriate persistence functions
	New(ctx context.Context, invoice interface{}) error
}
