package orders_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/razorpay/retail-store/internal/errors"
	"github.com/razorpay/retail-store/internal/orders"
)

func TestCoreImpl_GetOrderById(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockedRepo := orders.NewMockIRepo(ctrl)
	orders.SetRepo(mockedRepo)

	expectedOrderResp := orders.OrderResp{
		Id:     "testID",
		Status: "order placed",
		Value:  20,
	}
	orderResp := orders.Order{
		Id:         "testID",
		CustomerId: "cust_testID",
		ProductId:  "prod_testID",
		Quantity:   2,
		Status:     "order placed",
		Value:      20,
	}

	mockedRepo.EXPECT().GetOrderByIdFromDB("testID").Return(&orderResp, nil).Times(1)

	orders.NewCore()
	resp, err := orders.Core().GetOrderById("testID")
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, resp, &expectedOrderResp)
}

func TestCoreImpl_GetOrderById_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockedRepo := orders.NewMockIRepo(ctrl)
	orders.SetRepo(mockedRepo)

	mockedRepo.EXPECT().GetOrderByIdFromDB("testID").Return(nil, &errors.ErrorData{
		Code:    400,
		Message: "some error",
	}).Times(1)

	orders.NewCore()
	resp, err := orders.Core().GetOrderById("testID")
	assert.NotNil(t, err)
	assert.Nil(t, resp)
}

func TestCoreImpl_UpdateOrderStatus(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockedRepo := orders.NewMockIRepo(ctrl)
	orders.SetRepo(mockedRepo)

	mockedRepo.EXPECT().UpdateOrderStatusByIdInDB("testID", "failed").Return(&orders.Order{
		Id:     "testID",
		Status: "failed",
		Value:  20,
	}, nil).Times(1)

	orders.NewCore()
	resp, err := orders.Core().UpdateOrderStatus("testID", "failed")
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "failed", resp.Status)
}

func TestCoreImpl_UpdateOrderStatus_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockedRepo := orders.NewMockIRepo(ctrl)
	orders.SetRepo(mockedRepo)

	mockedRepo.EXPECT().UpdateOrderStatusByIdInDB("testID", "failed").Return(nil, &errors.ErrorData{
		Code:    400,
		Message: "some error",
	}).Times(1)

	orders.NewCore()
	resp, err := orders.Core().UpdateOrderStatus("testID", "failed")
	assert.NotNil(t, err)
	assert.Nil(t, resp)
}
