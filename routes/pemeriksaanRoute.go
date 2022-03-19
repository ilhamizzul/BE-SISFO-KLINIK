package routes

import (
	"BE-SISFO-KLINIK/controllers"

	"github.com/labstack/echo/v4"
)

func RoutePemeriksaan(e *echo.Echo) *echo.Echo {
	e.POST("/api/pemeriksaan/add", controllers.AddPemeriksaan)
	e.GET("/api/pemeriksaan/pasien", controllers.GetAllPemeriksaanForPasien)
	e.DELETE("/api/pemeriksaan/hapus", controllers.DeletePemerikaanById)
	e.POST("api/pemeriksaan/edit", controllers.EditPemeriksaanPasien)
	return e
}
