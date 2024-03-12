package products

import (
	"github.com/razorpay/retail-store/database"
	"github.com/razorpay/retail-store/internal/errors"
	"github.com/razorpay/retail-store/internal/uniqueId"
)

type IRepo interface {
	CreateProductInDB(product *Product) (*CreateProductResp, *errors.ErrorData)
	GetProductsFromDB() ([]Product, *errors.ErrorData)
	GetProductByIdFromDB(id string) (*Product, *errors.ErrorData)
	UpdateProductInDB(id string, input UpdateProductInput) (*Product, *errors.ErrorData)
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

func (r RepoImpl) CreateProductInDB(product *Product) (*CreateProductResp, *errors.ErrorData) {
	id, _ := uniqueId.New()
	product.SetID(id)
	if err := database.DB.Create(product).Error; err != nil {
		return nil, &errors.ErrorData{
			Code:    500,
			Message: err.Error(),
		}
	}

	var resp = CreateProductResp{
		Id:       product.Id,
		Name:     product.Name,
		Price:    product.Price,
		Quantity: product.Quantity,
		Message:  "product successfully added",
	}
	return &resp, nil
}

func (r RepoImpl) GetProductsFromDB() ([]Product, *errors.ErrorData) {
	var products []Product
	if err := database.DB.Find(&products).Error; err != nil {
		return nil, &errors.ErrorData{
			Code:    404,
			Message: err.Error(),
		}
	}
	return products, nil
}

func (r RepoImpl) GetProductByIdFromDB(id string) (*Product, *errors.ErrorData) {
	var product Product
	if err := database.DB.Where("id = ?", id).First(&product).Error; err != nil {
		return nil, &errors.ErrorData{
			Code:    404,
			Message: err.Error(),
		}
	}
	return &product, nil
}

func (r RepoImpl) UpdateProductInDB(id string, input UpdateProductInput) (*Product, *errors.ErrorData) {
	var product Product
	if err := database.DB.Where("id = ?", id).First(&product).Error; err != nil {
		return nil, &errors.ErrorData{
			Code:    404,
			Message: err.Error(),
		}
	}

	database.DB.Model(&product).Updates(input)
	return &product, nil
}
