package products_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/razorpay/retail-store/internal/products"
)

func TestCoreImpl_CreateProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockedRepo := products.NewMockIRepo(ctrl)
	products.SetRepo(mockedRepo)

	prodReq := products.Product{
		Name:     "testProd",
		Price:    10,
		Quantity: 100,
	}

	prodResp := products.CreateProductResp{
		Id:       "testId",
		Name:     "testProd",
		Price:    10,
		Quantity: 100,
		Message:  "product successfully added",
	}
	mockedRepo.EXPECT().CreateProductInDB(&prodReq).Return(&prodResp, nil).Times(1)

	products.NewCore()
	resp, err := products.Core().CreateProduct(&prodReq)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, resp, &prodResp)
}

func TestCoreImpl_GetProducts(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockedRepo := products.NewMockIRepo(ctrl)
	products.SetRepo(mockedRepo)

	prodResp := products.Product{
		Id:       "testId",
		Name:     "testProd",
		Price:    10,
		Quantity: 100,
	}
	mockedRepo.EXPECT().GetProductsFromDB().Return([]products.Product{prodResp}, nil).Times(1)

	products.NewCore()
	resp, err := products.Core().GetProducts()
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, resp[0], prodResp)
}

func TestCoreImpl_GetProductById(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockedRepo := products.NewMockIRepo(ctrl)
	products.SetRepo(mockedRepo)

	prodResp := products.Product{
		Id:       "testId",
		Name:     "testProd",
		Price:    10,
		Quantity: 100,
	}
	mockedRepo.EXPECT().GetProductByIdFromDB("testId").Return(&prodResp, nil).Times(1)

	products.NewCore()
	resp, err := products.Core().GetProductById("testId")
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, resp, &prodResp)
}

func TestCoreImpl_UpdateProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockedRepo := products.NewMockIRepo(ctrl)
	products.SetRepo(mockedRepo)

	input := products.UpdateProductInput{
		Price:    12,
		Quantity: 100,
	}
	prodResp := products.Product{
		Id:       "testId",
		Name:     "testProd",
		Price:    12,
		Quantity: 100,
	}
	mockedRepo.EXPECT().UpdateProductInDB("testId", input).Return(&prodResp, nil).Times(1)

	products.NewCore()
	resp, err := products.Core().UpdateProduct("testId", input)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, resp, &prodResp)
}
