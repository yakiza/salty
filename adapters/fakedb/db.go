package fakedb

import "github.com/yakiza/salty"

// Collection of all the repositories for the given database i.e. customer, invoices etc. contains higher level database
//functionality such as close()

type DB struct{}

func (DB) Implementation() string { return "fakeDB" }

func (DB) Customers() salty.CustomerRepository { return NewCustomerRepository() }
