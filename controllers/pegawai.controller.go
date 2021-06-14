package controllers

import (
	"net/http"

	"github.com/falahlaz/echo-tutorial/models"
	"github.com/labstack/echo"
)

func FetchAllPegawai(c echo.Context) error {
	result, err := models.FetchAllPegawai()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
