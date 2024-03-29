package customers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/razorpay/retail-store/internal/common"
)

func Create(c *gin.Context) {
	var customer *Customer
	_ = c.BindJSON(&customer)
	customer, err := Core().CreateCustomer(customer)
	if err != nil {
		code := err.GetHttpCode()
		c.JSON(code, err.Public())
		return
	}

	c.JSON(http.StatusOK, customer)
}

func GetMany(c *gin.Context) {
	customers, err := Core().GetCustomers()
	if err != nil {
		code := err.GetHttpCode()
		c.JSON(code, err.Public())
		return
	}

	c.JSON(http.StatusOK, GetManyResp{Items: customers})
}

func GetByID(c *gin.Context) {
	id := c.Params.ByName("id")
	customer, err := Core().GetCustomerById(id)
	if err != nil {
		code := err.GetHttpCode()
		c.JSON(code, err.Public())
		return
	}

	customer.Id = fmt.Sprintf("%s%s", common.CustomerIdPrefix, customer.Id)
	c.JSON(http.StatusOK, customer)
}

type GetManyResp struct {
	Items []Customer `json:"items"`
}
