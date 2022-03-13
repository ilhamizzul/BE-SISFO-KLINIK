package controllers

import (
	"BE-SISFO-KLINIK/pkg"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AddPemeriksaan(c echo.Context) error {

	return c.JSON(http.StatusOK, pkg.ResponseSuccess(200, true, "Add Pemeriksaan Successfully ", nil))
}

func GetAllPemeriksaanForPasien(c echo.Context) error {

	return c.JSON(http.StatusOK, pkg.ResponseSuccess(200, true, "Get Detail Pemeriksaan Successfully ", nil))
}
