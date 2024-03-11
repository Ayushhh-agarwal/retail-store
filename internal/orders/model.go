package orders

type Order struct {
	Id         string `json:"id"`
	CustomerId string `json:"customer_id"`
	ProductId  string `json:"product_id"`
	Quantity   int32  `json:"quantity"`
	Status     string `json:"status"`
	Value      int32  `json:"value"`
}

type OrderRequest struct {
	CustomerId      string           `json:"customer_id"`
	ProductsOrdered []ProductOrdered `json:"products_ordered"`
}

type ProductOrdered struct {
	ProductId string `json:"product_id"`
	Quantity  int32  `json:"quantity"`
}

type OrderResp struct {
	Id     string `json:"id"`
	Status string `json:"status"`
	Value  int32  `json:"value"`
}

type UpdateStatusRequest struct {
	Status string `json:"status"`
}

func (o *Order) TableName() string {
	return "orders"
}

func (o *Order) SetID(id string) {
	o.Id = id
}
