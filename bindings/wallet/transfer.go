package wallet

import "thunes/bindings"

type TransferReq struct {
	To     int `json:"to"`
	Amount int `json:"amount"`
}

type TransferRsp struct {
	Wallet *bindings.Wallet `json:"wallet"`
}
