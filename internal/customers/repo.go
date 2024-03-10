package customers

import (
	"github.com/razorpay/retail-store/database"
	"github.com/razorpay/retail-store/internal/errors"
	"github.com/razorpay/retail-store/internal/uniqueId"
)

func CreateCustomerInDB(customer *Customer) (error *errors.ErrorData) {
	id, _ := uniqueId.New()
	customer.SetID(id)
	if err := database.DB.Create(customer).Error; err != nil {
		return &errors.ErrorData{
			Code:    500,
			Message: err.Error(),
		}
	}
	return nil
}

func GetCustomersFromDB() ([]Customer, *errors.ErrorData) {
	var customers []Customer
	if err := database.DB.Find(&customers).Error; err != nil {
		return nil, &errors.ErrorData{
			Code:    404,
			Message: err.Error(),
		}
	}
	return customers, nil
}

func GetCustomerByIdFromDB(id string) (*Customer, *errors.ErrorData) {
	var customer Customer
	if err := database.DB.Where("id = ?", id).First(&customer).Error; err != nil {
		return nil, &errors.ErrorData{
			Code:    404,
			Message: err.Error(),
		}
	}
	return &customer, nil
}
