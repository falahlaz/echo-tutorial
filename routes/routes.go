package routes

import (
	"net/http"

	"github.com/falahlaz/echo-tutorial/controllers"
	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	server := echo.New()

	server.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World")
	})

	server.GET("/pegawai", controllers.FetchAllPegawai)

	return server
}
