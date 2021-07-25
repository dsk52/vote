package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type helloDTO struct {
	Message string `json:message`
}

func Hello(c echo.Context) error {
	helloDTO := new(helloDTO)
	if err := c.Bind(helloDTO); err != nil {
		return err
	}

	helloDTO.Message = "hello, world"

	return c.JSON(http.StatusOK, helloDTO)
}
