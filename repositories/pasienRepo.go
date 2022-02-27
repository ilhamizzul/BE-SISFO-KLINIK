package repositories

import (
	"BE-SISFO-KLINIK/config"
	"BE-SISFO-KLINIK/models"
)

func CreatePasienDB(data *models.Pasien) (*models.Pasien, error) {
	pasien := data
	db, err := config.ConnectionDatabase()
	if err != nil {
		return nil, err
	}

	err = db.Create(pasien).Error
	if err != nil {
		return nil, err
	}
	return pasien, nil
}
