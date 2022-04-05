package controllers

import (
	"BE-SISFO-KLINIK/models"
	"BE-SISFO-KLINIK/pkg"
	"BE-SISFO-KLINIK/repositories"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func AddPemeriksaan(c echo.Context) error {
	data := &models.Pemeriksaan{}
	// bind data

	if err := c.Bind(&data); err != nil {
		//fmt.Println("-------------", data.Id, "-", data.HasilPemeriksaan, "-", data.DeleteStatus, "-", data.StatusTransaksi, "-", "-", "-", "-")
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(400, false, "Bind data Error!"))
	}
	// after bind next to validate data
	err := pkg.ValidateAddPemeriksaan(data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(400, false, "Missing fields or data not valid!"))
	}

	savedData, err := repositories.CreatePemeriksaanDB(data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(400, false, err.Error()))
	}
	return c.JSON(http.StatusOK, pkg.ResponseSuccess(200, true, "Add Pemeriksaan Successfully ", savedData))
}

func GetAllPemeriksaanForPasien(c echo.Context) error {
	id := c.QueryParam("id")
	n, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(400, false, "id not valid"))
	}
	resultsData, row, err := repositories.GetAllPemeriksaanByIdPasien(n)
	if err != nil {
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(400, false, err.Error()))
	}
	result := pkg.CreatePagenationPemeriksaan(resultsData, row)
	return c.JSON(http.StatusOK, pkg.ResponseSuccess(200, true, "Get Detail Pemeriksaan Successfully ", result))
}

func DeletePemerikaanById(c echo.Context) error {
	id := c.QueryParam("id")
	n, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(400, false, "id not valid"))
	}
	_, status, err := repositories.DeletePemeriksaanDB(n)
	if err != nil {
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(400, false, err.Error()))
	} else if !status {
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(400, false, "Data Already Updated!"))
	}
	return c.JSON(http.StatusOK, pkg.ResponseSuccess(200, true, "Delete Pemeriksaan Successfully ", map[string]interface{}{"id": n}))
}

func EditPemeriksaanPasien(c echo.Context) error {
	data := &models.Pemeriksaan{}
	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(400, false, "Bind data Error!"))
	}
	PemeriksaanUpdated, numData, err := repositories.EditPemeriksaanPasienDB(data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(400, false, err.Error()))
	} else if !numData {
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(400, false, "Pasien Already Active!"))
	}
	return c.JSON(http.StatusOK, pkg.ResponseSuccess(200, true, "Activated Pasien Successfully ", PemeriksaanUpdated))
}
