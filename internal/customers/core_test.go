package customers_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/razorpay/retail-store/internal/customers"
	"github.com/razorpay/retail-store/internal/errors"
)

func TestCoreImpl_CreateCustomer(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockedRepo := customers.NewMockIRepo(ctrl)
	customers.SetRepo(mockedRepo)
	customer := customers.Customer{
		Name:    "testName",
		Phone:   "9876543223",
		Address: "testAddress",
		Email:   "testEmail",
	}
	mockedRepo.EXPECT().CreateCustomerInDB(&customer).Return(nil).Times(1)

	customers.NewCore()
	resp, err := customers.Core().CreateCustomer(&customer)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
}

func TestCoreImpl_CreateCustomer_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockedRepo := customers.NewMockIRepo(ctrl)
	customers.SetRepo(mockedRepo)
	customer := customers.Customer{
		Name:    "testName",
		Phone:   "9876543223",
		Address: "testAddress",
		Email:   "testEmail",
	}
	mockedRepo.EXPECT().CreateCustomerInDB(&customer).Return(&errors.ErrorData{Code: 400, Message: "some error"}).Times(1)

	customers.NewCore()
	resp, err := customers.Core().CreateCustomer(&customer)
	assert.NotNil(t, err)
	assert.Nil(t, resp)
	assert.Equal(t, err.Message, "some error")
}

func TestCoreImpl_GetCustomers(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockedRepo := customers.NewMockIRepo(ctrl)
	customers.SetRepo(mockedRepo)
	customer := customers.Customer{
		Name:    "testName",
		Phone:   "9876543223",
		Address: "testAddress",
		Email:   "testEmail",
	}
	mockedRepo.EXPECT().GetCustomersFromDB().Return([]customers.Customer{customer}, nil).Times(1)

	customers.NewCore()
	resp, err := customers.Core().GetCustomers()
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Len(t, resp, 1)
}

func TestCoreImpl_GetCustomerById(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockedRepo := customers.NewMockIRepo(ctrl)
	customers.SetRepo(mockedRepo)
	customer := customers.Customer{
		Id:      "testId",
		Name:    "testName",
		Phone:   "9876543223",
		Address: "testAddress",
		Email:   "testEmail",
	}
	mockedRepo.EXPECT().GetCustomerByIdFromDB("testId").Return(&customer, nil).Times(1)

	customers.NewCore()
	resp, err := customers.Core().GetCustomerById("testId")
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, resp.Id, "testId")
}
