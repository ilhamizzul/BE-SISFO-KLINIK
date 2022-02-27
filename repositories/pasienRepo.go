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

func GetAllPasienDB() ([]models.Pasien, int64, error) {
	var pasien []models.Pasien
	db, err := config.ConnectionDatabase()
	if err != nil {
		return nil, 0, err
	}
	result := db.Where("delete_status = ?", "false").Find(&pasien)
	if result.Error != nil {
		return nil, 0, err
	}

	return pasien, result.RowsAffected, nil
}

func SearchPasien(nama string) ([]models.Pasien, int64, error) {
	var pasien []models.Pasien
	db, err := config.ConnectionDatabase()
	if err != nil {
		return nil, 0, err
	}
	//result := db.Where("nama_pasien = ?", nama).Find(&pasien)
	result := db.Raw("SELECT * FROM pasiens WHERE MATCH( pasiens.nama_pasien ) AGAINST( ? )", nama).Scan(&pasien)
	if result.Error != nil {
		return nil, 0, err
	}

	return pasien, result.RowsAffected, nil
}
