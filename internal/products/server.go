package products

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/razorpay/retail-store/internal/common"
)

func Create(c *gin.Context) {
	var product *Product
	_ = c.BindJSON(&product)
	resp, err := CreateProduct(product)
	if err != nil {
		code := err.GetHttpCode()
		c.JSON(code, err.Public())
		return
	}

	c.JSON(http.StatusOK, resp)
}

func GetMany(c *gin.Context) {
	products, err := GetProducts()
	if err != nil {
		code := err.GetHttpCode()
		c.JSON(code, err.Public())
		return
	}

	c.JSON(http.StatusOK, GetManyResp{Items: products})
}

func GetByID(c *gin.Context) {
	id := c.Params.ByName("id")
	product, err := GetProductById(id)
	if err != nil {
		code := err.GetHttpCode()
		c.JSON(code, err.Public())
		return
	}

	product.Id = fmt.Sprintf("%s%s", common.ProductIdPrefix, product.Id)
	c.JSON(http.StatusOK, product)
}

func Update(c *gin.Context) {
	var input UpdateProductInput
	id := c.Params.ByName("id")

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	product, err := UpdateProduct(id, input)
	if err != nil {
		code := err.GetHttpCode()
		c.JSON(code, err.Public())
		c.Abort()
	}

	product.Id = fmt.Sprintf("%s%s", common.ProductIdPrefix, product.Id)
	c.JSON(http.StatusOK, product)
}

type GetManyResp struct {
	Items []Product `json:"items"`
}
