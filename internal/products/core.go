package products

import (
	"github.com/razorpay/retail-store/internal/errors"
)

func CreateProduct(product *Product) (*Product, *errors.ErrorData) {
	err := CreateProductInDB(product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func GetProducts() ([]Product, *errors.ErrorData) {
	products, err := GetProductsFromDB()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func GetProductById(id string) (*Product, *errors.ErrorData) {
	product, err := GetProductByIdFromDB(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func UpdateProduct(id string, input UpdateProductInput) (*Product, *errors.ErrorData) {
	product, err := UpdateProductInDB(id, input)
	if err != nil {
		return nil, err
	}
	return product, nil
}
