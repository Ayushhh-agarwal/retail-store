package transactions_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/razorpay/retail-store/internal/orders"
	"github.com/razorpay/retail-store/internal/transactions"
)

func TestCoreImpl_CreateTransaction(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockedRepo := transactions.NewMockIRepo(ctrl)
	mockedOrderCore := orders.NewMockICore(ctrl)
	transactions.SetRepo(mockedRepo)
	orders.SetCore(mockedOrderCore)

	txnReq := transactions.TransactionRequest{
		OrderId:       "ORD_testId",
		Value:         200,
		ModeOfPayment: "cash",
	}

	txnResp := transactions.Transaction{
		Id:            "TXN_testId",
		OrderId:       "ORD_testId",
		Value:         200,
		Status:        "initiated",
		ModeOfPayment: "cash",
	}

	updatedTxnResp := transactions.Transaction{
		Id:            "TXN_testId",
		OrderId:       "ORD_testId",
		Value:         200,
		Status:        "initiated",
		ModeOfPayment: "cash",
	}

	mockedOrderCore.EXPECT().GetOrderById("testId").Return(&orders.OrderResp{}, nil).Times(1)
	mockedRepo.EXPECT().CreateTransactionInDB(&txnReq).Return(&txnResp, nil).Times(1)

	mockedRepo.EXPECT().UpdateTransactionStatusByIdInDB("TXN_testId", "failed").Return(&updatedTxnResp, nil).Times(1)

	mockedOrderCore.EXPECT().UpdateOrderStatus("ORD_testId", "failed").Return(&orders.OrderResp{}, nil).Times(1)

	transactions.NewCore()
	resp, err := transactions.Core().CreateTransaction(&txnReq)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, resp, &updatedTxnResp)
}

func TestCoreImpl_GetTransactionStatusById(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockedRepo := transactions.NewMockIRepo(ctrl)
	transactions.SetRepo(mockedRepo)

	mockedRepo.EXPECT().GetTransactionStatusByIdFromDB("testId").Return("initiated", nil).Times(1)
	transactions.NewCore()
	resp, err := transactions.Core().GetTransactionStatusById("testId")
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, resp, "initiated")
}
