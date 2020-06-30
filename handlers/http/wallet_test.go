package http

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"thunes/bindings/wallet"
	"thunes/objects/business"
	"thunes/tools"
)

func TestWalletHandler_Get(t *testing.T) {
	// Setup
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/wallet/get", strings.NewReader("{}"))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	tools.AttachTokenInfo(c, &business.TokenInfo{
		Username: TestUsername,
		UID:      TestUID,
	})
	h := new(WalletHandler)

	// Assertions
	if assert.NoError(t, h.Get(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		rsp := new(struct {
			Code string        `json:"code"`
			Msg  string        `json:"msg"`
			Data wallet.GetRsp `json:"data"`
		})
		assert.NoError(t, json.NewDecoder(rec.Body).Decode(rsp))
		assert.Equal(t, "A0000", rsp.Code)
		assert.Equal(t, TestUID, rsp.Data.Wallet.UID)
		assert.Equal(t, TestSGD, rsp.Data.Wallet.SGD)
	}
}

func TestWalletHandler_Transfer(t *testing.T) {
	// Setup
	e := echo.New()
	amount := 10
	reqBody, _ := json.Marshal(&wallet.TransferReq{
		To:     TestTransferToUID,
		Amount: amount,
	})

	req := httptest.NewRequest(http.MethodPost, "/wallet/transfer", bytes.NewReader(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	tools.AttachTokenInfo(c, &business.TokenInfo{
		Username: TestUsername,
		UID:      TestUID,
	})
	h := new(WalletHandler)

	// Assertions
	if assert.NoError(t, h.Transfer(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		rsp := new(struct {
			Code string             `json:"code"`
			Msg  string             `json:"msg"`
			Data wallet.TransferRsp `json:"data"`
		})
		assert.NoError(t, json.NewDecoder(rec.Body).Decode(rsp))
		assert.Equal(t, "A0000", rsp.Code)
		assert.Equal(t, TestUID, rsp.Data.Wallet.UID)
		assert.Equal(t, TestSGD-amount, rsp.Data.Wallet.SGD)
	}
}

func TestWalletHandler_GetHistories(t *testing.T) {
	// Setup
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/wallet/get_history", strings.NewReader("{}"))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	tools.AttachTokenInfo(c, &business.TokenInfo{
		Username: TestUsername,
		UID:      TestUID,
	})
	h := new(WalletHandler)

	// Assertions
	if assert.NoError(t, h.GetHistories(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		rsp := new(struct {
			Code string                 `json:"code"`
			Msg  string                 `json:"msg"`
			Data wallet.GetHistoriesRsp `json:"data"`
		})
		assert.NoError(t, json.NewDecoder(rec.Body).Decode(rsp))
		assert.Equal(t, "A0000", rsp.Code, rsp.Msg)
		assert.Equal(t, 2, len(rsp.Data.Histories))
	}
}
