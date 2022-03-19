package repositories

import (
	"BE-SISFO-KLINIK/config"
	"BE-SISFO-KLINIK/models"
	"time"
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
	result := db.Where("id_pasien = ? AND delete_status = 0", id).Find(&pemeriksaan)
	if result.Error != nil {
		return nil, 0, err
	}

	return pemeriksaan, result.RowsAffected, nil
}

func DeletePemeriksaanDB(id int64) (models.Pemeriksaan, bool, error) {
	var pemeriksaan models.Pemeriksaan
	db, err := config.ConnectionDatabase()
	if err != nil {
		return pemeriksaan, false, err
	}
	result := db.Model(&pemeriksaan).Where("id = ? AND delete_status = 0", id).Updates(map[string]interface{}{"delete_status": true, "updated_at": time.Now()})
	if result.Error != nil {
		return pemeriksaan, false, result.Error
	} else if result.RowsAffected == 0 {
		return pemeriksaan, false, nil
	}
	return pemeriksaan, true, nil
}

func EditPemeriksaanPasienDB(data *models.Pemeriksaan) (*models.Pemeriksaan, bool, error) {
	db, err := config.ConnectionDatabase()
	if err != nil {
		return data, false, err
	}
	updateData := map[string]interface{}{
		"hasil_pemeriksaan": data.HasilPemeriksaan,
		"diagnosis":         data.Diagnosis,
		"terapi":            data.Terapi,
		"updated_at":        time.Now(),
	}
	result := db.Model(&data).Where("id = ? AND delete_status = 0", data.Id).Updates(updateData)
	if result.Error != nil {
		return data, false, result.Error
	} else if result.RowsAffected == 0 {
		return data, false, nil
	}
	return data, true, nil
}
