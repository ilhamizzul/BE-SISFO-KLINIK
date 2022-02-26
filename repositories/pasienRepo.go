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
func GetAllPasienDB() (*[]models.Pasien, error) {
	var pasien *[]models.Pasien
	db, err := config.ConnectionDatabase()
	if err != nil {
		return nil, err
	}

	result := db.Where("delete_status = ?", "false").Find(&pasien)
	if result.Error != nil {
		return nil, err
	}

	return pasien, nil
}
