package customers

import (
	"github.com/razorpay/retail-store/database"
	"github.com/razorpay/retail-store/internal/errors"
	"github.com/razorpay/retail-store/internal/uniqueId"
)

type IRepo interface {
	CreateCustomerInDB(customer *Customer) (error *errors.ErrorData)
	GetCustomersFromDB() ([]Customer, *errors.ErrorData)
	GetCustomerByIdFromDB(id string) (*Customer, *errors.ErrorData)
}

var repo IRepo

func NewRepo() IRepo {
	repo = &RepoImpl{}
	return repo
}

func Repo() IRepo {
	return repo
}

// SetRepo Used for setting up Mock IRepo
func SetRepo(r IRepo) {
	repo = r
}

type RepoImpl struct{}

func (r RepoImpl) CreateCustomerInDB(customer *Customer) (error *errors.ErrorData) {
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

func (r RepoImpl) GetCustomersFromDB() ([]Customer, *errors.ErrorData) {
	var customers []Customer
	if err := database.DB.Find(&customers).Error; err != nil {
		return nil, &errors.ErrorData{
			Code:    404,
			Message: err.Error(),
		}
	}
	return customers, nil
}

func (r RepoImpl) GetCustomerByIdFromDB(id string) (*Customer, *errors.ErrorData) {
	var customer Customer
	if err := database.DB.Where("id = ?", id).First(&customer).Error; err != nil {
		return nil, &errors.ErrorData{
			Code:    404,
			Message: err.Error(),
		}
	}
	return &customer, nil
}
