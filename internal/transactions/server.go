package transactions

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/razorpay/retail-store/internal/common"
)

func Create(c *gin.Context) {
	var transactionReq *TransactionRequest
	_ = c.BindJSON(&transactionReq)
	transaction, err := Core().CreateTransaction(transactionReq)
	if err != nil {
		code := err.GetHttpCode()
		c.JSON(code, err.Public())
		return
	}

	transaction.Id = fmt.Sprintf("%s%s", common.TransactionPrefix, transaction.Id)
	transaction.OrderId = fmt.Sprintf("%s%s", common.OrderIDPrefix, transaction.OrderId)
	c.JSON(http.StatusOK, transaction)
}

func GetStatusByID(c *gin.Context) {
	id := c.Params.ByName("id")
	status, err := Core().GetTransactionStatusById(id)
	if err != nil {
		code := err.GetHttpCode()
		c.JSON(code, err.Public())
		return
	}

	statusResp := GetTransactionStatusResponse{
		Status: status,
	}

	c.JSON(http.StatusOK, statusResp)
}
