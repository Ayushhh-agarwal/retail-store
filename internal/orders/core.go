package orders

import (
	"context"
	"fmt"

	"github.com/razorpay/retail-store/internal/customers"
	"github.com/razorpay/retail-store/internal/errors"
	"github.com/razorpay/retail-store/internal/mutex"
	"github.com/razorpay/retail-store/internal/products"
	"github.com/razorpay/retail-store/internal/uniqueId"
)

func CreateOrder(orderReq *OrderRequest) (*OrderResp, *errors.ErrorData) {
	valid, err := validateCreateOrderRequest(orderReq)
	order := Order{}
	id, _ := uniqueId.New()
	order.SetID(id)
	var value = 0
	if !valid {
		return nil, err
	}

	var productKeysForLock []string
	for _, prodDetails := range orderReq.ProductsOrdered {
		productKeysForLock = append(productKeysForLock, prodDetails.ProductId)
	}

	ctx := context.Background()
	_, mutexErr := mutex.AcquireAndReleaseMultiple(ctx, productKeysForLock, func(lockAcquiredFailedIds []string, _ []string) (interface{}, *errors.ErrorData) {
		if len(lockAcquiredFailedIds) > 0 {
			return nil, &errors.ErrorData{
				Code:    400,
				Message: "unable to take lock. Another operation is in progress.",
			}
		}

		for _, productOrderDetail := range orderReq.ProductsOrdered {
			product, err := validateProductDetails(productOrderDetail)
			if err != nil {
				return nil, err
			}

			productId := productOrderDetail.ProductId[5:]
			value = value + int(product.Price*productOrderDetail.Quantity)
			products.UpdateProduct(productId, products.UpdateProductInput{
				Quantity: product.Quantity - productOrderDetail.Quantity,
			})
		}
		return nil, nil
	})
	if mutexErr != nil {
		return nil, &errors.ErrorData{
			Code:    400,
			Message: errors.RedisLockError,
		}
	}

	err = CreateOrderInDB(&order)
	if err != nil {
		return nil, err
	}
	orderResp := OrderResp{
		Id:     order.Id,
		Status: order.Status,
		Value:  int32(value),
	}
	return &orderResp, nil
}

func GetOrderById(id string) (*OrderResp, *errors.ErrorData) {
	order, err := GetOrderByIdFromDB(id)
	if err != nil {
		return nil, err
	}

	orderResp := OrderResp{
		Id:     order.Id,
		Status: order.Status,
	}
	return &orderResp, nil
}

func UpdateOrderStatus(id, status string) (*OrderResp, *errors.ErrorData) {
	order, err := UpdateOrderStatusByIdInDB(id, status)
	if err != nil {
		return nil, err
	}

	orderResp := OrderResp{
		Id:     order.Id,
		Status: order.Status,
		Value:  order.Value,
	}
	return &orderResp, nil
}

func validateCreateOrderRequest(request *OrderRequest) (bool, *errors.ErrorData) {
	valid := validateCustomerId(request.CustomerId)
	if !valid {
		return false, &errors.ErrorData{
			Code:    400,
			Message: "invalid customer id, please create customer",
		}
	}

	return true, nil
}

func validateCustomerId(customerId string) bool {
	customerId = customerId[4:]
	_, err := customers.GetCustomerById(customerId)
	if err != nil {
		return false
	}
	return true
}

func validateProductDetails(productOrderDetail ProductOrdered) (*products.Product, *errors.ErrorData) {
	productId := productOrderDetail.ProductId[5:]
	product, err := products.GetProductById(productId)
	if err != nil {
		return nil, err
	}

	if product.Quantity < productOrderDetail.Quantity {
		return nil, &errors.ErrorData{
			Code:    400,
			Message: fmt.Sprintf("Ordered quantity of %s is more than available quantity", product.Name),
		}
	}

	return product, nil
}
