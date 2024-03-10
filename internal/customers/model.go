package customers

type Customer struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	Email   string `json:"email"`
}

func (c *Customer) TableName() string {
	return "customers"
}

func (c *Customer) SetID(id string) {
	c.Id = id
}
