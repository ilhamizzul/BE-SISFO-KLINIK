package repositories

import (
	"BE-SISFO-KLINIK/config"
	"BE-SISFO-KLINIK/models"
	"time"
)

func CreateObatDB(data *models.Obat) (*models.Obat, error) {
	obat := data
	db, err := config.ConnectionDatabase()
	if err != nil {
		return nil, err
	}

	err = db.Create(obat).Error
	if err != nil {
		return nil, err
	}
	return obat, nil
}

type ResultAllObats struct {
	Id        int64
	Kode      string
	Nama      string
	HargaJual int64
	Masuk     int64
	Keluar    int64
	Sisa      int64
}

var Obat *models.Obat

func GetAllObatDB() ([]ResultAllObats, error) {
	var result []ResultAllObats
	db, err := config.ConnectionDatabase()
	if err != nil {
		return nil, err
	}
	// err = db.Model(obat).Preload(string(clause.)).Find(obat).Error
	//r := db.Raw("SELECT a.id,a.kode,a.nama,a.harga_jual,b.masuk,b.keluar,b.sisa FROM obats a JOIN stok_obats b ON a.id = b.id_obat;").Scan(&result)
	r := db.Raw("SELECT a.id,a.kode,a.nama,a.harga_jual,b.masuk,b.keluar,b.sisa FROM obats a JOIN stok_obats b ON a.id = b.id_obat WHERE ( YEAR(a.updated_at) = YEAR(curdate()) ) AND ( MONTH(a.updated_at) = MONTH(curdate())  ) AND (a.delete_status = 0) ORDER BY (b.sisa) DESC").Scan(&result)
	if r.Error != nil {
		return nil, r.Error
	}
	return result, nil
}

func GetAllTrashObatDB() ([]ResultAllObats, error) {
	var result []ResultAllObats
	db, err := config.ConnectionDatabase()
	if err != nil {
		return nil, err
	}
	// err = db.Model(obat).Preload(string(clause.)).Find(obat).Error
	//r := db.Raw("SELECT a.id,a.kode,a.nama,a.harga_jual,b.masuk,b.keluar,b.sisa FROM obats a JOIN stok_obats b ON a.id = b.id_obat;").Scan(&result)
	r := db.Raw("SELECT a.id,a.kode,a.nama,a.harga_jual,b.masuk,b.keluar,b.sisa FROM obats a JOIN stok_obats b ON a.id = b.id_obat WHERE ( YEAR(a.updated_at) = YEAR(curdate()) ) AND ( MONTH(a.updated_at) = MONTH(curdate())  ) AND (a.delete_status = 1) ORDER BY (b.sisa) DESC").Scan(&result)
	if r.Error != nil {
		return nil, r.Error
	}
	return result, nil
}

func DeleteObatDB(idObat int64) (*models.Obat, error, bool) {
	var obat *models.Obat
	db, err := config.ConnectionDatabase()
	if err != nil {
		return nil, nil, false
	}
	result := db.Model(&obat).Where("id = ? AND delete_status = 0", idObat).Updates(map[string]interface{}{"delete_status": true, "updated_at": time.Now()})
	if result.Error != nil {
		return obat, result.Error, false
	} else if result.RowsAffected == 0 {
		return obat, nil, false
	}
	return obat, nil, true

}

func ActivateObatDB(idObat int64) (*models.Obat, error, bool) {
	var obat *models.Obat
	db, err := config.ConnectionDatabase()
	if err != nil {
		return nil, nil, false
	}
	result := db.Model(&obat).Where("id = ? AND delete_status = 1", idObat).Updates(map[string]interface{}{"delete_status": false, "updated_at": time.Now()})
	if result.Error != nil {
		return obat, result.Error, false
	} else if result.RowsAffected == 0 {
		return obat, nil, false
	}
	return obat, nil, true

}

func EditObatDB(data *models.UpdateObat) (*models.UpdateObat, bool, error) {
	var obat *models.Obat
	db, err := config.ConnectionDatabase()
	if err != nil {
		return data, false, err
	}
	updateData := map[string]interface{}{
		"nama":       data.Nama,
		"harga_jual": data.HargaJual,
		"updated_at": time.Now(),
	}
	result := db.Model(&obat).Where("id = ? AND delete_status = 0", data.Id).Updates(updateData)
	if result.Error != nil {
		return data, false, result.Error
	} else if result.RowsAffected == 0 {
		return data, false, nil
	}
	return data, true, nil
}
