package base

type Transaction struct {
	Id        int `json:"id"`
	IdService int `json:"service"`
	IdOrder   int `json:"order"`
	Amount    int `json:"amount" binding:"required,gte=0"`
	Flag      int `json:"flag" binding:"required,gte=0, lte=2"`
}

type Transactions []Transaction
