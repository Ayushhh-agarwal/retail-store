package customers

import (
	"github.com/razorpay/retail-store/internal/errors"
)

type ICore interface {
	CreateCustomer(customer *Customer) (*Customer, *errors.ErrorData)
	GetCustomers() ([]Customer, *errors.ErrorData)
	GetCustomerById(id string) (*Customer, *errors.ErrorData)
}

var core ICore

type CoreImpl struct{}

func NewCore() ICore {
	core = &CoreImpl{}
	return core
}

func SetCore(c ICore) {
	core = c
}

func Core() ICore {
	return core
}

func (c CoreImpl) CreateCustomer(customer *Customer) (*Customer, *errors.ErrorData) {
	err := Repo().CreateCustomerInDB(customer)
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func (c CoreImpl) GetCustomers() ([]Customer, *errors.ErrorData) {
	customers, err := Repo().GetCustomersFromDB()
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (c CoreImpl) GetCustomerById(id string) (*Customer, *errors.ErrorData) {
	customer, err := Repo().GetCustomerByIdFromDB(id)
	if err != nil {
		return nil, err
	}
	return customer, nil
}
