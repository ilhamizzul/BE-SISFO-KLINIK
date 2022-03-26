package controllers

import (
	"BE-SISFO-KLINIK/models"
	"BE-SISFO-KLINIK/pkg"
	"BE-SISFO-KLINIK/repositories"
	"net/http"

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
	err = repositories.CreateStokDB(savedData.Id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(400, false, err.Error()))
	}

	return c.JSON(http.StatusOK, pkg.ResponseSuccess(200, true, "Add Obat Successfully ", savedData))
}

func GetAllObat(c echo.Context) error {
	Data, err := repositories.GetAllObatDB()

	if err != nil {
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(400, false, err.Error()))
	}
	return c.JSON(http.StatusOK, pkg.ResponseSuccess(200, true, "Get All Obat Successfully ", Data))
}
