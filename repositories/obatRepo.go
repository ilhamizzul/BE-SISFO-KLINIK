package repositories

import (
	"BE-SISFO-KLINIK/config"
	"BE-SISFO-KLINIK/models"
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
