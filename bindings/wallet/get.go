package wallet

type GetRsp struct {
	Wallet *Wallet `json:"wallet"`
}

type Wallet struct {
	UID int `json:"uid"`
	SGD int `json:"sgd"`
}
