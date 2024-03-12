package transactions

import (
	"github.com/razorpay/retail-store/database"
	"github.com/razorpay/retail-store/internal/errors"
	"github.com/razorpay/retail-store/internal/transactions/transactionStatusTypes"
	"github.com/razorpay/retail-store/internal/uniqueId"
)

type IRepo interface {
	CreateTransactionInDB(txnReq *TransactionRequest) (*Transaction, *errors.ErrorData)
	GetTransactionStatusByIdFromDB(id string) (string, *errors.ErrorData)
	UpdateTransactionStatusByIdInDB(id, status string) (*Transaction, *errors.ErrorData)
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

func (r RepoImpl) CreateTransactionInDB(txnReq *TransactionRequest) (*Transaction, *errors.ErrorData) {
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

func (r RepoImpl) GetTransactionStatusByIdFromDB(id string) (string, *errors.ErrorData) {
	var txn Transaction
	if err := database.DB.Where("id = ?", id).First(&txn).Error; err != nil {
		return "", &errors.ErrorData{
			Code:    404,
			Message: err.Error(),
		}
	}
	return txn.Status, nil
}

func (r RepoImpl) UpdateTransactionStatusByIdInDB(id, status string) (*Transaction, *errors.ErrorData) {
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
