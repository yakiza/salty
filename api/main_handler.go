package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/yakiza/salty/api/internal"
	"github.com/yakiza/salty/internal/customer"
	"net/http"
)

func NewHandler() http.Handler {
	r := chi.NewRouter()

	useCase := customer.NewCreateCustomerUseCase()

	r.Mount(internal.HealthCheckEndPoint, healthCheckController())
	r.Mount(internal.CustomerEndPoint, NewCustomerController(useCase))

	return r
}
