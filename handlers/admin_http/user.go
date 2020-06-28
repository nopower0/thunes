package admin_http

import (
	"github.com/labstack/echo/v4"
	"thunes/bindings"
	"thunes/bindings/user"
	"thunes/components"
	"thunes/objects/models"
	"thunes/settings"
	"thunes/tools"
	"time"
)

type UserHandler struct {
}

func (h *UserHandler) Login(c echo.Context) error {
	tokenInfo := tools.GetTokenInfo(c)
	if tokenInfo.UID != 0 {
		return bindings.JSONResponse(c, bindings.UserAlreadyLoginError)
	}

	req := new(user.LoginReq)
	if err := c.Bind(req); err != nil {
		return bindings.JSONResponse(c, bindings.NewParamError(err.Error()))
	}
	if len(req.Username) == 0 || len(req.Password) == 0 {
		return bindings.JSONResponse(c, bindings.InvalidUsernameOrPasswordError)
	}

	passwordHash := tools.PasswordHash(req.Password)

	if u, err := models.DefaultUserManager.GetByCredential(req.Username, passwordHash); err != nil {
		return err
	} else if u == nil {
		return bindings.JSONResponse(c, bindings.InvalidUsernameOrPasswordError)
	} else {
		tokenInfo.UID = u.UID
		tokenInfo.Username = u.Username
		if err := components.DefaultAuthService.UpdateTokenInfo(
			c.Request().Header.Get(settings.HeaderToken),
			tokenInfo,
			time.Now().Add(settings.TokenTTL),
		); err != nil {
			return err
		}
		return bindings.JSONResponse(c, &user.LoginRsp{
			UID:      u.UID,
			Username: u.Username,
		})
	}
}
