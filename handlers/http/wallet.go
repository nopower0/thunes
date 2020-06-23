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
	rspWallet := &bindings.Wallet{
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

func (*WalletHandler) Transfer(c echo.Context) error {
	req := new(wallet.TransferReq)
	if err := c.Bind(req); err != nil {
		return bindings.JSONResponse(c, bindings.NewParamError(err.Error()))
	}

	tokenInfo := tools.GetTokenInfo(c)
	// Check transfer to user
	if req.To == tokenInfo.UID {
		return bindings.JSONResponse(c, bindings.NewParamError("Cannot transfer to self"))
	} else if u, err := models.DefaultUserManager.Get(req.To); err != nil {
		return err
	} else if u == nil {
		return bindings.JSONResponse(c, bindings.TransferToNotExistError)
	}
	// Check amount
	if req.Amount <= 0 {
		return bindings.JSONResponse(c, bindings.NewParamError("Invalid amount"))
	}

	if receipt, err := models.DefaultWalletManager.Transfer(tokenInfo.UID, req.To, req.Amount); err != nil {
		if err == models.InsufficientBalanceError {
			return bindings.JSONResponse(c, bindings.InsufficientBalanceError)
		}
		return err
	} else {
		w := &bindings.Wallet{
			UID: receipt.From.UID,
			SGD: receipt.From.SGD,
		}
		return bindings.JSONResponse(c, &wallet.TransferRsp{
			Wallet: w,
		})
	}
}
