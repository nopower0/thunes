package wallet

type GetHistoriesReq struct {
	Start  int `json:"start"`
	Length int `json:"length"`
}

type GetHistoriesRsp struct {
	Histories []*TransferHistory `json:"histories"`
}

type TransferHistory struct {
	To              *User `json:"to"`
	Amount          int   `json:"amount"`
	TransactionTime int   `json:"transaction_time"`
}

type User struct {
	UID      int    `json:"uid"`
	Username string `json:"username"`
}
