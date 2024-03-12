package transactions

import (
	"math/rand"

	"github.com/razorpay/retail-store/internal/errors"
	"github.com/razorpay/retail-store/internal/orders"
	"github.com/razorpay/retail-store/internal/orders/orderStatusTypes"
	"github.com/razorpay/retail-store/internal/transactions/transactionStatusTypes"
)

type ICore interface {
	CreateTransaction(transactionReq *TransactionRequest) (*Transaction, *errors.ErrorData)
	GetTransactionStatusById(id string) (string, *errors.ErrorData)
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

func (c CoreImpl) CreateTransaction(transactionReq *TransactionRequest) (*Transaction, *errors.ErrorData) {
	valid, err := validateCreateTransactionRequest(transactionReq)
	if !valid {
		return nil, err
	}

	txn, err := Repo().CreateTransactionInDB(transactionReq)
	if err != nil {
		return nil, err
	}

	// Assuming API Call is done to 3rd Party Payment gateway
	// expecting 2 results success or failed from the API
	// so making random response from the apiResult slice
	apiResult := []string{"success", "failed"}
	random := rand.Intn(2)

	status := transactionStatusTypes.Completed
	orderStatus := orderStatusTypes.Processed
	if apiResult[random] == "failed" {
		status = transactionStatusTypes.Failed
		orderStatus = orderStatusTypes.Failed
	}

	txn, err = Repo().UpdateTransactionStatusByIdInDB(txn.Id, status)
	if err != nil {
		return nil, err
	}
	_, err = orders.Core().UpdateOrderStatus(txn.OrderId, orderStatus)
	if err != nil {
		return nil, err
	}
	return txn, nil
}

func (c CoreImpl) GetTransactionStatusById(id string) (string, *errors.ErrorData) {
	status, err := Repo().GetTransactionStatusByIdFromDB(id)
	if err != nil {
		return "", err
	}

	return status, nil
}

func validateCreateTransactionRequest(request *TransactionRequest) (bool, *errors.ErrorData) {
	valid := validateOrderId(request.OrderId)
	if !valid {
		return false, &errors.ErrorData{
			Code:    400,
			Message: "invalid customer id, please create customer",
		}
	}

	return true, nil
}

func validateOrderId(orderId string) bool {
	orderId = orderId[4:]
	_, err := orders.Core().GetOrderById(orderId)
	if err != nil {
		return false
	}
	return true
}
