package products

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Price    int32  `json:"price"`
	Quantity int32  `json:"quantity"`
}

type UpdateProductInput struct {
	Price    int32 `json:"price"`
	Quantity int32 `json:"quantity"`
}

type CreateProductResp struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Price    int32  `json:"price"`
	Quantity int32  `json:"quantity"`
	Message  string `json:"message"`
}

func (p *Product) TableName() string {
	return "products"
}

func (p *Product) SetID(id string) {
	p.Id = id
}
