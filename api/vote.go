package api

import (
	"net/http"

	"github.com/dsk52/vote/db"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type MVote struct {
	gorm.Model
	Name string
}

type VoteListDTO struct {
	List []MVote `json: list`
}

func VoteList(c echo.Context) error {
	d := db.NewDB()
	db := d.Connect()

	votes := []MVote{}
	db.Find(&votes)

	vl := new(VoteListDTO)
	vl.List = votes

	return c.JSON(http.StatusOK, vl)
}

func VoteSave(c echo.Context) error {
	return c.String(http.StatusOK, "vote")
}
