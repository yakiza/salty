package internal

import (
	"github.com/yakiza/salty/internal"
	"strconv"
)

type Customer struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
}

func MarshallCustomer(customer Customer) (internal.Customer, error) {
	mobile, err := strconv.Atoi(customer.Mobile)
	if err != nil {
		return internal.Customer{}, err
	}

	return internal.Customer{
		Name:     internal.CustomerName(customer.Name),
		LastName: internal.CustomerLastName(customer.LastName),
		Email:    internal.CustomerEmail(customer.Email),
		Mobile:   internal.CustomerMobile(mobile),
	}, nil
}
