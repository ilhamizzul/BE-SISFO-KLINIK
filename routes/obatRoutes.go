package routes

import (
	"BE-SISFO-KLINIK/controllers"

	"github.com/labstack/echo/v4"
)

func RouteObat(e *echo.Echo) *echo.Echo {
	e.POST("/api/obat", controllers.CreateObat)
	e.GET("/api/obat", controllers.GetAllObat)
	return e
}
