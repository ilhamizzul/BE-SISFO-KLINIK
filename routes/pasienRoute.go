package routes

import (
	"BE-SISFO-KLINIK/controllers"

	"github.com/labstack/echo/v4"
)

func RoutePasien(e *echo.Echo) *echo.Echo {
	e.POST("/api/pasien/add", controllers.AddPasien)
	return e
}
