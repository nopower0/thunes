package http

import (
	"github.com/labstack/echo/v4"
	"thunes/bindings"
	"thunes/bindings/wallet"
	"thunes/objects/models"
	"thunes/tools"
)

type WalletHandler struct {
}

func (*WalletHandler) Get(c echo.Context) error {
	tokenInfo := tools.GetTokenInfo(c)
	rspWallet := &wallet.Wallet{
		UID: tokenInfo.UID,
	}

	if w, err := models.DefaultWalletManager.Get(tokenInfo.UID); err != nil {
		return err
	} else if w != nil {
		rspWallet.SGD = w.SGD
	}

	return bindings.JSONResponse(c, &wallet.GetRsp{
		Wallet: rspWallet,
	})
}
