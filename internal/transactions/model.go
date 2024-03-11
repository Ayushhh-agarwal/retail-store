package transactions

type Transaction struct {
	Id            string `json:"id"`
	OrderId       string `json:"order_id"`
	Value         int32  `json:"value"`
	Status        string `json:"status"`
	ModeOfPayment string `json:"mode_of_payment"`
}

type TransactionRequest struct {
	OrderId       string `json:"order_id"`
	Value         int32  `json:"value"`
	ModeOfPayment string `json:"mode_of_payment"`
}

type UpdateTransactionRequest struct {
	Status string `json:"status"`
}

type GetTransactionStatusResponse struct {
	Status string `json:"status"`
}

func (t *Transaction) TableName() string {
	return "transactions"
}

func (t *Transaction) SetID(id string) {
	t.Id = id
}
