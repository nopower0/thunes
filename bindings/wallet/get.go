package wallet

import "thunes/bindings"

type GetRsp struct {
	Wallet *bindings.Wallet `json:"wallet"`
}
