package customers

type Customer struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	Email   string `json:"email"`
}

func (c *Customer) TableName() string {
	return "customers"
}
