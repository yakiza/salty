package api

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/yakiza/salty"
	"github.com/yakiza/salty/api/internal"
	"net/http"
)

type CustomerController struct {
	chi.Router
	UseCase salty.CustomerUseCases
}

func (h CustomerController) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var customer internal.Customer
	err := json.NewDecoder(r.Body).Decode(&customer)

	marshalledCustomer, err := internal.MarshallCustomer(customer)
	if err != nil {
		http.Error(w, fmt.Sprintf("malformed customer data, %v", err), http.StatusBadRequest)
	}

	err = h.UseCase.Create(r.Context(), marshalledCustomer)
	if err != nil {
		http.Error(w, fmt.Sprintf("error creting customer, %v", err), http.StatusBadRequest)
	}

}

func NewCustomerController(useCase salty.CustomerUseCases) CustomerController {
	c := CustomerController{
		chi.NewRouter(),
		useCase,
	}

	c.Post("/create", c.CreateCustomer)

	return c
}
