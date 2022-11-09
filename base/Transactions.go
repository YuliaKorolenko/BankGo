package base

type Transaction struct {
	Id        int `json:"id"`
	IdService int `json:"service"`
	IdOrder   int `json:"order"`
	Amount    int `json:"amount" binding:"required,gte=0"`
}

type Transactions []Transaction
