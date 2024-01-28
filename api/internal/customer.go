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
		internal.CustomerName(customer.Name),
		internal.CustomerLastName(customer.LastName),
		internal.CustoemrEmail(customer.Email),
		internal.CustomerMobile(mobile),
	}, nil
}
