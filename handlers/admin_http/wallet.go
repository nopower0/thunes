package admin_http

import (
	"github.com/labstack/echo/v4"
	"thunes/bindings"
	"thunes/objects/models"
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
