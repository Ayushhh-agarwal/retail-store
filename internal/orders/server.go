package orders

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/razorpay/retail-store/internal/common"
)

func Create(c *gin.Context) {
	var orderReq *OrderRequest
	_ = c.BindJSON(&orderReq)
	orderResp, err := CreateOrder(orderReq)
	if err != nil {
		code := err.GetHttpCode()
		c.JSON(code, err.Public())
		return
	}

	orderResp.Id = fmt.Sprintf("%s%s", common.OrderIDPrefix, orderResp.Id)
	c.JSON(http.StatusOK, orderResp)
}

func GetByID(c *gin.Context) {
	id := c.Params.ByName("id")
	orderResp, err := GetOrderById(id)
	if err != nil {
		code := err.GetHttpCode()
		c.JSON(code, err.Public())
		return
	}

	orderResp.Id = fmt.Sprintf("%s%s", common.OrderIDPrefix, orderResp.Id)
	c.JSON(http.StatusOK, orderResp)
}
