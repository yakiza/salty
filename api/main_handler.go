package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/yakiza/salty/api/internal"
	"net/http"
)

func NewHandler() http.Handler {
	r := chi.NewRouter()

	r.Mount(internal.HealthCheckEndPoint, healthCheckController())

	return r
}
