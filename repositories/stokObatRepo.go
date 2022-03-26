package repositories

import (
	"BE-SISFO-KLINIK/config"
	"BE-SISFO-KLINIK/models"
	"time"
)

func CreateStokDB(id int64) error {
	StokObat := &models.StokObat{
		IdObat: id,
		Date:   time.Now(),
	}

	db, err := config.ConnectionDatabase()
	if err != nil {
		return err
	}

	err = db.Create(StokObat).Error
	if err != nil {
		return err
	}
	return nil
}
