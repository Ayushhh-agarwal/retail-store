package products

import (
	"github.com/razorpay/retail-store/database"
	"github.com/razorpay/retail-store/internal/errors"
	"github.com/razorpay/retail-store/internal/uniqueId"
)

func CreateProductInDB(product *Product) (error *errors.ErrorData) {
	id, _ := uniqueId.New()
	product.SetID(id)
	if err := database.DB.Create(product).Error; err != nil {
		return &errors.ErrorData{
			Code:    500,
			Message: err.Error(),
		}
	}
	return nil
}

func GetProductsFromDB() ([]Product, *errors.ErrorData) {
	var products []Product
	if err := database.DB.Find(&products).Error; err != nil {
		return nil, &errors.ErrorData{
			Code:    404,
			Message: err.Error(),
		}
	}
	return products, nil
}

func GetProductByIdFromDB(id string) (*Product, *errors.ErrorData) {
	var product Product
	if err := database.DB.Where("id = ?", id).First(&product).Error; err != nil {
		return nil, &errors.ErrorData{
			Code:    404,
			Message: err.Error(),
		}
	}
	return &product, nil
}

func UpdateProductInDB(id string, input UpdateProductInput) (*Product, *errors.ErrorData) {
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
