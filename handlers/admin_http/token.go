package admin_http

import (
	"github.com/labstack/echo/v4"
	"thunes/bindings"
	"thunes/bindings/token"
	"thunes/components"
	"thunes/objects/business"
	"thunes/objects/models"
	"thunes/settings"
	"thunes/tools"
	"time"
)

type TokenHandler struct{}

func (*TokenHandler) Request(c echo.Context) error {
	req := new(token.RequestReq)
	if err := c.Bind(req); err != nil {
		return bindings.JSONResponse(c, bindings.NewParamError(err.Error()))
	}

	tokenInfo := new(business.TokenInfo)
	expireAt := time.Now().Add(settings.TokenTTL)

	if len(req.Username) != 0 || len(req.Password) != 0 {
		passwordHash := tools.PasswordHash(req.Password)
		if user, err := models.DefaultUserManager.GetByCredential(req.Username, passwordHash); err != nil {
			return err
		} else if user == nil {
			return bindings.JSONResponse(c, bindings.InvalidUsernameOrPasswordError)
		} else {
			tokenInfo.Username = user.Username
			tokenInfo.UID = user.UID
		}
	}

	if t, err := components.DefaultAuthService.CreateTokenInfo(tokenInfo, expireAt); err != nil {
		return err
	} else {
		return bindings.JSONResponse(c, &token.RequestRsp{
			Token:    t,
			ExpireAt: int(expireAt.Unix()),
		})
	}
}
