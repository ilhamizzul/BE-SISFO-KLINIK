package pkg

import (
	"BE-SISFO-KLINIK/models"

	"github.com/go-playground/validator"
)

func ValidateAddPasien(pasien *models.Pasien) error {
	validate := validator.New()
	err := validate.Struct(pasien)
	return err
}

func ValidateAddPemeriksaan(pemeriksaan *models.Pemeriksaan) error {
	validate := validator.New()
	err := validate.Struct(pemeriksaan)
	return err
}
