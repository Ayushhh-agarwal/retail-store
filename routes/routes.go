package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/razorpay/retail-store/internal/status_check"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/status", status_check.Get)
	return r
}