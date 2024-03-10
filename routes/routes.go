package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/razorpay/retail-store/internal/customers"
	"github.com/razorpay/retail-store/internal/products"
	"github.com/razorpay/retail-store/internal/status_check"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/status", status_check.Get)

	grpCustomer := r.Group("/customer")
	{
		grpCustomer.POST("", customers.Create)
		grpCustomer.GET("", customers.GetMany)
		grpCustomer.GET(":id", customers.GetByID)
	}

	grpProduct := r.Group("/products")
	{
		grpProduct.POST("", products.Create)
		grpProduct.GET("", products.GetMany)
		grpProduct.GET(":id", products.GetByID)
		grpProduct.PATCH(":id", products.Update)
	}
	return r
}
