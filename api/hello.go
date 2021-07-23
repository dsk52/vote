package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HelloDTO struct {
	Message string `json:message`
}

func Hello(c echo.Context) error {
	HelloDTO := new(HelloDTO)
	if err := c.Bind(HelloDTO); err != nil {
		return err
	}

	HelloDTO.Message = "aaa"

	return c.JSON(http.StatusOK, HelloDTO)
}
