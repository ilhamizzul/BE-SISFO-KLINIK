package controllers

import (
	"BE-SISFO-KLINIK/models"
	"BE-SISFO-KLINIK/pkg"
	"BE-SISFO-KLINIK/repositories"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateObat(c echo.Context) error {
	data := &models.Obat{}
	// bind data
	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(400, false, "Bind data Error!"))
	}
	// after bind next to validate data
	err := pkg.ValidateAddObat(data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(400, false, "Missing fields or data not valid!"))
	}

	savedData, err := repositories.CreateObatDB(data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(400, false, err.Error()))
	}
	// err = repositories.CreateStokDB(savedData.Id)
	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, pkg.ResponseError(400, false, err.Error()))
	// }

	return c.JSON(http.StatusOK, pkg.ResponseSuccess(200, true, "Add Obat Successfully ", savedData))
}

func GetAllObat(c echo.Context) error {
	Data, err := repositories.GetAllObatDB()
	if err != nil {
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(400, false, err.Error()))
	}
	result := pkg.CreatePagenationObat(Data, len(Data))
	return c.JSON(http.StatusOK, pkg.ResponseSuccess(200, true, "Get All Obat Successfully ", result))
}

func GetAllTrashObat(c echo.Context) error {
	Data, err := repositories.GetAllTrashObatDB()
	if err != nil {
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(400, false, err.Error()))
	}
	result := pkg.CreatePagenationObat(Data, len(Data))
	return c.JSON(http.StatusOK, pkg.ResponseSuccess(200, true, "Get All Trash Obat Successfully ", result))
}

func DeactiveObat(c echo.Context) error {
	id := c.Param("id")
	idObat, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(400, false, "id not valid"))
	}
	deletedData, err, numData := repositories.DeleteObatDB(idObat)
	if err != nil {
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(400, false, err.Error()))
	} else if !numData {
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(400, false, "Data Already Deleted!"))
	}
	return c.JSON(http.StatusOK, pkg.ResponseSuccess(200, true, "Delete Obat Successfully", deletedData))
}

func ActivatedObat(c echo.Context) error {
	id := c.Param("id")
	idObat, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(400, false, "id not valid"))
	}
	deletedData, err, numData := repositories.ActivateObatDB(idObat)
	if err != nil {
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(400, false, err.Error()))
	} else if !numData {
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(400, false, "Data Already Active!"))
	}
	return c.JSON(http.StatusOK, pkg.ResponseSuccess(200, true, "Activated Obat Successfully", deletedData))
}

func EditObat(c echo.Context) error {
	return c.JSON(http.StatusOK, pkg.ResponseSuccess(200, true, "Delete Obat Successfully", nil))
}
