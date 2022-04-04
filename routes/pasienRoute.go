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
	e.GET("/api/pasien/sampah", controllers.GetAllPasienByName)
	e.POST("/api/pasien/aktivasi", controllers.ActivatePasien)
	e.POST("/api/pasien/edit", controllers.EditPasien)
	e.GET("/api/pasien/detail", controllers.GetAllPasienById)
	return e
}
