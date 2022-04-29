package routes

import (
	"BE-SISFO-KLINIK/controllers"

	"github.com/labstack/echo/v4"
)

func RouteObat(e *echo.Echo) *echo.Echo {
	e.POST("/api/obat", controllers.CreateObat)
	e.GET("/api/obat", controllers.GetAllObat)
	e.GET("/api/obat/sampah", controllers.GetAllTrashObat)
	e.DELETE("/api/obat/deactive/:id", controllers.DeactiveObat)
	e.PATCH("/api/obat/activated/:id", controllers.ActivatedObat)
	e.PATCH("/api/obat", controllers.EditObat)
	return e
}
