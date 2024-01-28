package internal

type Customer struct {
	Name     CustomerName
	LastName CustomerLastName
	Email    CustoemrEmail
	Mobile   CustomerMobile
}

type CustomerName string
type CustomerLastName string
type CustoemrEmail string
type CustomerMobile int
