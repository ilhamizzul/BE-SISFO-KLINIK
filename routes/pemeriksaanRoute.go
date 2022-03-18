package routes

import (
	"BE-SISFO-KLINIK/controllers"

	"github.com/labstack/echo/v4"
)

func RoutePemeriksaan(e *echo.Echo) *echo.Echo {
	e.POST("/api/pemeriksaan/add", controllers.AddPemeriksaan)
	e.GET("/api/pemeriksaan/pasien", controllers.GetAllPemeriksaanForPasien)
	return e
}
