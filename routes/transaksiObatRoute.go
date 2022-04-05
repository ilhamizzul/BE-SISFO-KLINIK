package routes

import (
	"BE-SISFO-KLINIK/controllers"

	"github.com/labstack/echo/v4"
)

func RouteTransaksi(e *echo.Echo) *echo.Echo {
	e.POST("/api/transaksi/:id_pemeriksaan", controllers.AddTransaksiObat)
	return e
}
