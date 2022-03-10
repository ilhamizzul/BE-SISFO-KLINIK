package repositories

import (
	"BE-SISFO-KLINIK/config"
	"BE-SISFO-KLINIK/models"
	"time"
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
	result := db.Where("delete_status = ?", false).Find(&pasien)
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
		return nil, 0, result.Error
	}

	return pasien, result.RowsAffected, nil
}

func DeletePasienDB(id int64) (models.Pasien, error, bool) {
	var pasien models.Pasien
	db, err := config.ConnectionDatabase()
	if err != nil {
		return pasien, err, false
	}

	result := db.Model(&pasien).Where("id_pemeriksaan = ? AND delete_status = 0", id).Updates(map[string]interface{}{"delete_status": true, "updated_at": time.Now()})
	if result.Error != nil {
		return pasien, result.Error, false
	} else if result.RowsAffected == 0 {
		return pasien, nil, false
	}
	return pasien, nil, true
}

func ActivatePasienDB(id int64) (models.Pasien, error, bool) {
	var pasien models.Pasien
	db, err := config.ConnectionDatabase()
	if err != nil {
		return pasien, err, false
	}

	result := db.Model(&pasien).Where("id_pemeriksaan = ? AND delete_status = 1", id).Updates(map[string]interface{}{"delete_status": false, "updated_at": time.Now()})
	if result.Error != nil {
		return pasien, result.Error, false
	} else if result.RowsAffected == 0 {
		return pasien, nil, false
	}
	return pasien, nil, true
}

func GetAllPasienDeleteDB() ([]models.Pasien, int64, error) {
	var pasien []models.Pasien
	db, err := config.ConnectionDatabase()
	if err != nil {
		return nil, 0, err
	}
	result := db.Where("delete_status = ?", true).Find(&pasien)
	if result.Error != nil {
		return nil, 0, err
	}

	return pasien, result.RowsAffected, nil
}
