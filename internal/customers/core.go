package customers

import (
	"github.com/razorpay/retail-store/internal/errors"
)

func CreateCustomer(customer *Customer) (*Customer, *errors.ErrorData) {
	err := CreateCustomerInDB(customer)
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func GetCustomers() ([]Customer, *errors.ErrorData) {
	customers, err := GetCustomersFromDB()
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func GetCustomerById(id string) (*Customer, *errors.ErrorData) {
	customer, err := GetCustomerByIdFromDB(id)
	if err != nil {
		return nil, err
	}
	return customer, nil
}
