package orders

import (
	"github.com/razorpay/retail-store/database"
	"github.com/razorpay/retail-store/internal/errors"
	"github.com/razorpay/retail-store/internal/orders/orderStatusTypes"
	"github.com/razorpay/retail-store/internal/uniqueId"
)

func CreateOrderInDB(order *Order) (error *errors.ErrorData) {
	id, _ := uniqueId.New()
	order.SetID(id)
	order.Status = orderStatusTypes.OrderPlaced
	if err := database.DB.Create(order).Error; err != nil {
		return &errors.ErrorData{
			Code:    500,
			Message: err.Error(),
		}
	}
	return nil
}

func GetOrderByIdFromDB(id string) (*Order, *errors.ErrorData) {
	var order Order
	if err := database.DB.Where("id = ?", id).First(&order).Error; err != nil {
		return nil, &errors.ErrorData{
			Code:    404,
			Message: err.Error(),
		}
	}
	return &order, nil
}

func UpdateOrderStatusByIdInDB(id, status string) (*Order, *errors.ErrorData) {
	var order Order
	input := UpdateStatusRequest{
		Status: status,
	}
	if err := database.DB.Where("id = ?", id).First(&order).Error; err != nil {
		return nil, &errors.ErrorData{
			Code:    404,
			Message: err.Error(),
		}
	}

	database.DB.Model(&order).Updates(input)
	return &order, nil
}
