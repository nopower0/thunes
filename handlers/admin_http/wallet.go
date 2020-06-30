package admin_http

import (
	"github.com/labstack/echo/v4"
	"thunes/bindings"
	"thunes/objects/models"
	"thunes/tools"
	"time"
)

type WalletHandler struct {
}

func (*WalletHandler) GetSummary(c echo.Context) error {
	if summary, err := models.DefaultWalletAnalysisManager.GetWalletSummary(); err != nil {
		return err
	} else {
		record := new(struct {
			TotalUser int `json:"total_user"`
			TotalSGD  int `json:"total_sgd"`
		})
		record.TotalUser = summary.TotalUser
		record.TotalSGD = summary.TotalSGD
		return bindings.JSONResponse(c, map[string]interface{}{
			"record": record,
		})
	}
}

func (*WalletHandler) GetTransactionSummary(c echo.Context) error {
	year, month, day := time.Now().Date()
	today := time.Date(year, month, day, 0, 0, 0, 0, time.Now().Location())
	nearestSunday := today.AddDate(0, 0, -int(today.Weekday()-time.Sunday))
	if summary, err := models.DefaultWalletAnalysisManager.GetTransactionSummary(nearestSunday.AddDate(0, 0, -7), nearestSunday); err != nil {
		return err
	} else {
		record := new(struct {
			Count    int `json:"count"`
			TotalSGD int `json:"total_sgd"`
		})
		record.Count = summary.Count
		record.TotalSGD = summary.TotalSGD
		return bindings.JSONResponse(c, map[string]interface{}{
			"last_week": record,
		})
	}
}

func (*WalletHandler) List(c echo.Context) error {
	wallets, err := models.DefaultWalletAnalysisManager.GetAllWallets()
	if err != nil {
		return err
	}
	uids := make([]int, len(wallets))
	for i, w := range wallets {
		uids[i] = w.UID
	}
	users, err := models.DefaultUserManager.GetMany(uids)
	if err != nil {
		return err
	}

	records := make([]*struct {
		UID      int    `json:"uid"`
		Username string `json:"username"`
		SGD      int    `json:"sgd"`
	}, len(wallets))
	for i, w := range wallets {
		u := users[w.UID]
		records[i] = &struct {
			UID      int    `json:"uid"`
			Username string `json:"username"`
			SGD      int    `json:"sgd"`
		}{
			UID:      w.UID,
			Username: u.Username,
			SGD:      w.SGD,
		}
	}
	return bindings.JSONResponse(c, map[string]interface{}{
		"records": records,
	})
}

func (*WalletHandler) Create(c echo.Context) error {
	req := new(struct {
		Username string `json:"username"`
		Password string `json:"password"`
		SGD      int    `json:"sgd"`
	})
	if err := c.Bind(req); err != nil {
		return bindings.JSONResponse(c, bindings.NewParamError(err.Error()))
	}
	if len(req.Username) == 0 || len(req.Password) == 0 {
		return bindings.JSONResponse(c, bindings.InvalidUsernameOrPasswordError)
	}

	passwordHash := tools.PasswordHash(req.Password)
	if u, err := models.DefaultUserManager.Create(req.Username, passwordHash); err != nil {
		return err
	} else if w, err := models.DefaultWalletManager.Create(u.UID, req.SGD); err != nil {
		return err
	} else {
		record := new(struct {
			UID      int    `json:"uid"`
			Username string `json:"username"`
			SGD      int    `json:"sgd"`
		})
		record.UID = u.UID
		record.Username = u.Username
		record.SGD = w.SGD
		return bindings.JSONResponse(c, map[string]interface{}{
			"record": record,
		})
	}
}
