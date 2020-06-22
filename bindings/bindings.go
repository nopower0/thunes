package bindings

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func JSONResponse(c echo.Context, i interface{}) error {
	if e, ok := i.(*Error); ok {
		return c.JSON(http.StatusOK, &struct {
			Code string      `json:"code"`
			Msg  string      `json:"msg"`
			Data interface{} `json:"data"`
		}{
			Code: e.Code,
			Msg:  e.Message,
			Data: nil,
		})
	} else {
		return c.JSON(http.StatusOK, &struct {
			Code string      `json:"code"`
			Msg  string      `json:"msg"`
			Data interface{} `json:"data"`
		}{
			Code: "A0000",
			Msg:  "Success",
			Data: i,
		})
	}
}
