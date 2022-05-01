package repositories

import (
	"BE-SISFO-KLINIK/config"
	"BE-SISFO-KLINIK/models"
	"time"

	"gorm.io/gorm"
)

var stokObats *models.StokObat
var Pemeriksaan *models.Pemeriksaan

func AddTransaksiObatDB(data []models.TransaksiObat, batch int, IdPemeriksaan int64) ([]models.TransaksiObat, error) {
	i := 0
	var StokObat models.StokObat
	db, err := config.ConnectionDatabase()
	if err != nil {
		return nil, err
	}
	statusTx := db.Transaction(func(tx *gorm.DB) error {
		tx.CreateInBatches(data, batch)
		updateData := map[string]interface{}{
			"status_transaksi":  "Sudah",
			"tanggal_beli_obat": time.Now(),
			"updated_at":        time.Now(),
		}
		for i = 0; i < batch; i++ {
			status := tx.Raw("SELECT * FROM db_klinik.stok_obats a WHERE id_obat = ? AND YEAR(a.created_at) = YEAR(curdate())  AND ( MONTH(a.created_at) = MONTH(curdate())  ) AND (a.delete_status = 0)", data[i].IdObat).Scan(&StokObat)
			if status.Error != nil {
				return status.Error
			}
			status = db.Model(&stokObats).Where("id = ? AND delete_status = 0", StokObat.Id).Updates(map[string]interface{}{"sisa": StokObat.Sisa - data[i].Jumlah, "keluar": StokObat.Keluar + data[i].Jumlah, "updated_at": time.Now()})
			if status.Error != nil {
				return status.Error
			}
		}
		status := tx.Model(&Pemeriksaan).Where("id=?", IdPemeriksaan).Updates(updateData)
		if status.Error != nil {
			return status.Error
		}
		return nil
	})
	if statusTx != nil {
		return nil, statusTx
	}
	return nil, nil
}

func GetAllTransactionsByPemeriksaanId(IdPemeriksaan int64, IdPasien int64) (interface{}, error) {
	var Pemeriksaan models.Pemeriksaan
	var TransaksiObat []models.TransaksiObat
	var Pasien models.Pasien
	var ArrObat []models.Obat
	var Obat models.Obat
	db, err := config.ConnectionDatabase()
	if err != nil {
		return nil, err
	}
	r := db.Where("id = ? AND id_pasien = ?", IdPemeriksaan, IdPasien).First(&Pemeriksaan)
	if r.Error != nil {
		return nil, err
	}
	if r.RowsAffected == 0 {
		return nil, nil
	}
	statusTx := db.Transaction(func(tx *gorm.DB) error {
		result := tx.Where("delete_status = ? AND id_pemeriksaan = ?", false, IdPemeriksaan).Find(&TransaksiObat)
		if result.Error != nil {
			return err
		}
		result = tx.Where("delete_status = ? AND id = ?", false, IdPasien).First(&Pasien)
		if result.Error != nil {
			return err
		}

		for _, item := range TransaksiObat {
			result = db.Where("delete_status = ? AND id = ?", false, item.IdObat).Scan(&Obat)
			if result.Error != nil {
				return err
			}
			ArrObat = append(ArrObat, Obat)
		}

		return nil
	})

	if statusTx != nil {
		return nil, statusTx
	}

	return map[string]interface{}{
		"pasien": Pasien,
		"obat": map[string]interface{}{
			"detail_transasksi": TransaksiObat,
			"detail_obat":       ArrObat,
		},
	}, nil
}
