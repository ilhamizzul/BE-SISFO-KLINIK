package repositories

import (
	"BE-SISFO-KLINIK/config"
	"BE-SISFO-KLINIK/models"
	"time"

	"gorm.io/gorm"
)

var Pemeriksaan *models.Pemeriksaan

func AddTransaksiObatDB(data []models.TransaksiObat, batch int, IdPemeriksaan int64) ([]models.TransaksiObat, error) {
	db, err := config.ConnectionDatabase()
	if err != nil {
		return nil, err
	}

	db.Transaction(func(tx *gorm.DB) error {
		tx.CreateInBatches(data, batch)
		updateData := map[string]interface{}{
			"status_transaksi":  "Sudah",
			"tanggal_beli_obat": time.Now(),
			"updated_at":        time.Now(),
		}
		tx.Model(&Pemeriksaan).Where("id=?", IdPemeriksaan).Updates(updateData)
		return nil
	})
	return nil, nil
}
