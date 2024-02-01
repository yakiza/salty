package api_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/suite"
	"github.com/yakiza/salty"
	"github.com/yakiza/salty/api"
	"github.com/yakiza/salty/internal"
	"net/http"
	"net/http/httptest"
	"testing"
)

var _ salty.CustomerUseCases = &CustomerTestDoubles{}

type CustomerTestDoubles struct {
	CreateTD func(ctx context.Context, customer internal.Customer) error
}

func (td CustomerTestDoubles) Create(ctx context.Context, customer internal.Customer) error {
	return td.CreateTD(ctx, customer)
}

type CustomerHandlerTestSuite struct {
	suite.Suite

	handler http.Handler

	useCase CustomerTestDoubles
}

func (suite *CustomerHandlerTestSuite) SetupSuite() {
	suite.useCase = CustomerTestDoubles{}
	suite.handler = api.NewHandler(&suite.useCase)
}

func (suite *CustomerHandlerTestSuite) TestCustomerCreateSuccess() {
	expectedCustomer := map[string]interface{}{
		"name":      "John",
		"last_name": "Doe",
		"email":     "john.doe@example.com",
		"mobile":    "445555555555",
	}
	suite.useCase.CreateTD = func(_ context.Context, incomingCustomer internal.Customer) error {
		// TODO: Assert equality of the incoming customer data  with the data we did send (expectedCustomer)
		return nil
	}

	body, err := json.Marshal(expectedCustomer)
	suite.Require().NoError(err)

	r, err := http.NewRequest(http.MethodPost, fmt.Sprintf("/customer/create"), bytes.NewReader(body))
	suite.Require().NoError(err)

	r.Header.Add("Content-Type", "application/json")
	w := httptest.NewRecorder()
	suite.handler.ServeHTTP(w, r)

	suite.Equal(http.StatusOK, w.Code)
	suite.Empty(w.Body.Bytes())
}

func TestCustomerTestSuite(t *testing.T) {
	suite.Run(t, new(CustomerHandlerTestSuite))
}
