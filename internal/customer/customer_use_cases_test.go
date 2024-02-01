package customer_test

import (
	"context"
	"github.com/stretchr/testify/suite"
	"github.com/yakiza/salty"
	"github.com/yakiza/salty/internal"
)

type CustomerPersistenceTD struct {
	CreateTD func(ctx context.Context, customer internal.Customer) error
}

func (td CustomerPersistenceTD) Create(ctx context.Context, customer internal.Customer) error {
	return td.CreateTD(ctx, customer)
}

type CustomerTestSuite struct {
	suite.Suite

	customerTD salty.CustomerRepository
}

func (suite *CustomerTestSuite) SetupSuite() {
}

func (suite *CustomerTestSuite) TestCustomerCreateSuccess() {

}
