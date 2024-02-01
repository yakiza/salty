package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/yakiza/salty"
	"github.com/yakiza/salty/api/internal"
	"net/http"
)

func NewHandler(customerUseCases salty.CustomerUseCases) http.Handler {
	r := chi.NewRouter()

	r.Mount(internal.HealthCheckEndPoint, healthCheckController())
	r.Mount(internal.CustomerEndPoint, NewCustomerHandler(customerUseCases))

	return r
}
