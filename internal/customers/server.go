package customers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var customer *Customer
	_ = c.BindJSON(&customer)
	customer, err := CreateCustomer(customer)
	if err != nil {
		code := err.GetHttpCode()
		c.JSON(code, err.Public())
		c.Abort()
	}

	c.JSON(http.StatusOK, customer)
}

func GetMany(c *gin.Context) {
	customers, err := GetCustomers()
	if err != nil {
		code := err.GetHttpCode()
		c.JSON(code, err.Public())
		c.Abort()
	}

	c.JSON(http.StatusOK, GetManyResp{Items: customers})
}

func GetByID(c *gin.Context) {
	id := c.Params.ByName("id")
	customer, err := GetCustomerById(id)
	if err != nil {
		code := err.GetHttpCode()
		c.JSON(code, err.Public())
		c.Abort()
	}

	c.JSON(http.StatusOK, customer)
}

type GetManyResp struct {
	Items []Customer `json:"items"`
}
