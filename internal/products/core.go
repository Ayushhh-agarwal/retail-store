package products

import (
	"github.com/razorpay/retail-store/internal/errors"
)

type ICore interface {
	CreateProduct(product *Product) (*CreateProductResp, *errors.ErrorData)
	GetProducts() ([]Product, *errors.ErrorData)
	GetProductById(id string) (*Product, *errors.ErrorData)
	UpdateProduct(id string, input UpdateProductInput) (*Product, *errors.ErrorData)
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

func (c CoreImpl) CreateProduct(product *Product) (*CreateProductResp, *errors.ErrorData) {
	resp, err := Repo().CreateProductInDB(product)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c CoreImpl) GetProducts() ([]Product, *errors.ErrorData) {
	products, err := Repo().GetProductsFromDB()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (c CoreImpl) GetProductById(id string) (*Product, *errors.ErrorData) {
	product, err := Repo().GetProductByIdFromDB(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (c CoreImpl) UpdateProduct(id string, input UpdateProductInput) (*Product, *errors.ErrorData) {
	product, err := Repo().UpdateProductInDB(id, input)
	if err != nil {
		return nil, err
	}
	return product, nil
}
