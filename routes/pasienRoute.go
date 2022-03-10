package routes

import (
	"BE-SISFO-KLINIK/controllers"

	"github.com/labstack/echo/v4"
)

func RoutePasien(e *echo.Echo) *echo.Echo {
	e.POST("/api/pasien/add", controllers.AddPasien)
	e.GET("/api/pasien", controllers.GetAllPasien)
	e.GET("/api/pasien/cari", controllers.SearchPasien)
	e.POST("/api/pasien/hapus", controllers.DeletePasien)
	e.GET("/api/pasien/sampah", controllers.GetAllPasienDelete)
	e.POST("/api/pasien/aktivasi", controllers.ActivatePasien)
	return e
}
