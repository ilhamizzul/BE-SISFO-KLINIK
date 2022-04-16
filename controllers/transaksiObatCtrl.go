package controllers

import (
	"BE-SISFO-KLINIK/models"
	"BE-SISFO-KLINIK/pkg"
	"BE-SISFO-KLINIK/repositories"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func AddTransaksiObat(c echo.Context) error {
	param := c.Param("id_pemeriksaan")
	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(400, false, "id not valid"))
	}
	// todo cek id_pemeriksaan ada atau engga!

	data := []models.TransaksiObat{}
	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(400, false, "Bind data Error!"))
	}
	fmt.Println("--------------------------------")
	if len(data) == 0 {
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(400, false, "Data Not Found"))
	}

	result, err := repositories.AddTransaksiObatDB(data, len(data), id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(400, false, err.Error()))
	}
	return c.JSON(http.StatusOK, pkg.ResponseSuccess(200, true, "Add Pasien Successfully ", result))
}
