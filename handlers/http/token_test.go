package http

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"thunes/bindings/token"
	"time"
)

func TestTokenHandler_Request(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/token/request", strings.NewReader("{}"))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := new(TokenHandler)

	// Assertions
	if assert.NoError(t, h.Request(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		rsp := new(struct {
			Code string           `json:"code"`
			Msg  string           `json:"msg"`
			Data token.RequestRsp `json:"data"`
		})
		assert.NoError(t, json.NewDecoder(rec.Body).Decode(rsp))
		assert.Equal(t, "A0000", rsp.Code)
		assert.Equal(t, TestTokenNotLogin, rsp.Data.Token)
		assert.Greater(t, rsp.Data.ExpireAt, int(time.Now().Unix()))
	}
}
