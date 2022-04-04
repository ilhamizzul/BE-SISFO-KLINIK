package controllers

import (
	"BE-SISFO-KLINIK/models"
	"BE-SISFO-KLINIK/pkg"
	"BE-SISFO-KLINIK/repositories"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func AddPasien(c echo.Context) error {
	data := &models.Pasien{}
	// bind data

	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(400, false, "Bind data Error!"))
	}
	// after bind next to validate data
	err := pkg.ValidateAddPasien(data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(400, false, "Missing fields or data not valid!"))
	}

	savedData, err := repositories.CreatePasienDB(data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(400, false, err.Error()))
	}

	return c.JSON(http.StatusOK, pkg.ResponseSuccess(200, true, "Add Pasien Successfully ", savedData))
}

func GetAllPasien(c echo.Context) error {
	resultsData, row, err := repositories.GetAllPasienDB()
	if err != nil {
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(400, false, err.Error()))
	}
	result := pkg.CreatePagenationV2(resultsData, row)
	return c.JSON(http.StatusOK, pkg.ResponseSuccess(200, true, "Get All Pasien Successfully ", result))
}

func GetAllPasienByName(c echo.Context) error {
	resultsData, row, err := repositories.GetAllPasienDeleteDB()
	if err != nil {
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(400, false, err.Error()))
	}
	result := pkg.CreatePagenationV2(resultsData, row)
	return c.JSON(http.StatusOK, pkg.ResponseSuccess(200, true, "Get All Pasien by Name Successfully ", result))
}

func SearchPasien(c echo.Context) error {
	nama := c.QueryParam("nama")
	resultsData, row, err := repositories.SearchPasien(nama)
	if err != nil {
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(400, false, err.Error()))
	}
	result := pkg.CreatePagenationV2(resultsData, row)
	return c.JSON(http.StatusOK, pkg.ResponseSuccess(200, true, "Get Pasien Successfully ", result))
}

func DeletePasien(c echo.Context) error {
	data := &models.Pasien{}
	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(400, false, "Bind data Error!"))
	}
	deletedData, err, numData := repositories.DeletePasienDB(data.Id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(400, false, err.Error()))
	} else if !numData {
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(400, false, "Data Already Updated!"))
	}
	return c.JSON(http.StatusOK, pkg.ResponseSuccess(200, true, "Delete Pasien Successfully ", map[string]interface{}{"id": deletedData.Id}))
}

func ActivatePasien(c echo.Context) error {
	data := &models.Pasien{}
	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(400, false, "Bind data Error!"))
	}
	ActivatedData, err, numData := repositories.ActivatePasienDB(data.Id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(400, false, err.Error()))
	} else if !numData {
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(400, false, "Pasien Already Active!"))
	}
	return c.JSON(http.StatusOK, pkg.ResponseSuccess(200, true, "Activated Pasien Successfully ", map[string]interface{}{"id": ActivatedData.Id}))
}

func EditPasien(c echo.Context) error {
	data := &models.Pasien{}
	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(400, false, "Bind data Error!"))
	}
	PasienUpdated, numData, err := repositories.EditPasienDB(data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(400, false, err.Error()))
	} else if !numData {
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(400, false, "Pasien Already Active!"))
	}
	return c.JSON(http.StatusOK, pkg.ResponseSuccess(200, true, "Activated Pasien Successfully ", PasienUpdated))
}

func GetAllPasienById(c echo.Context) error {
	id := c.QueryParam("id")
	n, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(400, false, "id not valid"))
	}
	resultsData, err := repositories.GetAllPasienByIdDB(n)
	if err != nil {
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(400, false, err.Error()))
	}
	return c.JSON(http.StatusOK, pkg.ResponseSuccess(200, true, "Get Pasien by id Successfully ", resultsData))
}
