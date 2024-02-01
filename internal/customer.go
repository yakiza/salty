package internal

type Customer struct {
	Name     CustomerName
	LastName CustomerLastName
	Email    CustomerEmail
	Mobile   CustomerMobile
}

type CustomerName string
type CustomerLastName string
type CustomerEmail string
type CustomerMobile int
