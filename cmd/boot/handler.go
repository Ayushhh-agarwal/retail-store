package boot

import (
	"github.com/razorpay/retail-store/internal/customers"
	"github.com/razorpay/retail-store/internal/orders"
	"github.com/razorpay/retail-store/internal/products"
	"github.com/razorpay/retail-store/internal/transactions"
)

func LoadAllModules() {

	customers.NewCore()
	customers.NewRepo()

	orders.NewCore()
	orders.NewRepo()

	products.NewCore()
	products.NewRepo()

	transactions.NewCore()
	transactions.NewRepo()
}
