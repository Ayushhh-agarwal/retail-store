package transactions

import (
	"github.com/razorpay/retail-store/database"
	"github.com/razorpay/retail-store/internal/errors"
	"github.com/razorpay/retail-store/internal/transactions/transactionStatusTypes"
	"github.com/razorpay/retail-store/internal/uniqueId"
)

func CreateTransactionInDB(txnReq *TransactionRequest) (*Transaction, *errors.ErrorData) {
	id, _ := uniqueId.New()
	txnReq.OrderId = txnReq.OrderId[4:]
	txn := Transaction{
		Id:            id,
		OrderId:       txnReq.OrderId,
		Value:         txnReq.Value,
		Status:        transactionStatusTypes.Initiated,
		ModeOfPayment: txnReq.ModeOfPayment,
	}

	if err := database.DB.Create(txn).Error; err != nil {
		return nil, &errors.ErrorData{
			Code:    500,
			Message: err.Error(),
		}
	}
	return &txn, nil
}

func GetTransactionStatusByIdFromDB(id string) (string, *errors.ErrorData) {
	var txn Transaction
	if err := database.DB.Where("id = ?", id).First(&txn).Error; err != nil {
		return "", &errors.ErrorData{
			Code:    404,
			Message: err.Error(),
		}
	}
	return txn.Status, nil
}

func UpdateTransactionStatusByIdInDB(id, status string) (*Transaction, *errors.ErrorData) {
	var txn Transaction
	input := UpdateTransactionRequest{
		Status: status,
	}

	if err := database.DB.Where("id = ?", id).First(&txn).Error; err != nil {
		return nil, &errors.ErrorData{
			Code:    404,
			Message: err.Error(),
		}
	}

	database.DB.Model(&txn).Updates(input)
	return &txn, nil
}
