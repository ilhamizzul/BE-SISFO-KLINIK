package routes

import (
	"BE-SISFO-KLINIK/controllers"

	"github.com/labstack/echo/v4"
)

func RouteTransaksi(e *echo.Echo) *echo.Echo {
	e.POST("/api/transaksi/:id_pemeriksaan", controllers.AddTransaksiObat)
	e.GET("/api/transaksi/:id_pemeriksaan/:id_pasien", controllers.GetTransaksiPasienByPemeriksaanId)
	return e
}
