package http

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"thunes/bindings"
	"thunes/bindings/user"
	"thunes/objects/business"
	"thunes/tools"
)

func TestUserHandler_Login(t *testing.T) {
	// Setup
	e := echo.New()
	reqBody, _ := json.Marshal(&user.LoginReq{
		Username: TestUsername,
		Password: TestPassword,
	})

	// Valid case
	{
		req := httptest.NewRequest(http.MethodPost, "/user/login", bytes.NewReader(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		tools.AttachTokenInfo(c, &business.TokenInfo{})
		h := new(UserHandler)

		// Assertions
		if assert.NoError(t, h.Login(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			rsp := new(struct {
				Code string        `json:"code"`
				Msg  string        `json:"msg"`
				Data user.LoginRsp `json:"data"`
			})
			assert.NoError(t, json.NewDecoder(rec.Body).Decode(rsp))
			assert.Equal(t, "A0000", rsp.Code, rsp.Msg)
			assert.Equal(t, TestUsername, rsp.Data.Username)
			assert.Equal(t, TestUID, rsp.Data.UID)
		}
	}

	// User already login case
	{
		req := httptest.NewRequest(http.MethodPost, "/user/login", bytes.NewReader(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		tools.AttachTokenInfo(c, &business.TokenInfo{
			Username: TestUsername,
			UID:      TestUID,
		})
		h := new(UserHandler)

		// Assertions
		if assert.NoError(t, h.Login(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			rsp := new(struct {
				Code string        `json:"code"`
				Msg  string        `json:"msg"`
				Data user.LoginRsp `json:"data"`
			})
			assert.NoError(t, json.NewDecoder(rec.Body).Decode(rsp))
			assert.Equal(t, bindings.UserAlreadyLoginError.Code, rsp.Code, rsp.Msg)
		}
	}

}
