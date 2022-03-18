package repositories

import (
	"BE-SISFO-KLINIK/config"
	"BE-SISFO-KLINIK/models"
)

func CreatePemeriksaanDB(data *models.Pemeriksaan) (*models.Pemeriksaan, error) {
	pemeriksaan := data
	db, err := config.ConnectionDatabase()
	if err != nil {
		return nil, err
	}

	err = db.Create(pemeriksaan).Error
	if err != nil {
		return nil, err
	}
	return pemeriksaan, nil
}

func GetAllPemeriksaanByIdPasien(id int64) ([]models.Pemeriksaan, int64, error) {
	var pemeriksaan []models.Pemeriksaan
	db, err := config.ConnectionDatabase()
	if err != nil {
		return nil, 0, err
	}
	result := db.Where("id_pasien = ?", id).Find(&pemeriksaan)
	if result.Error != nil {
		return nil, 0, err
	}

	return pemeriksaan, result.RowsAffected, nil
}
