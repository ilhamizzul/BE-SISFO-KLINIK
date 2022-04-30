package config

import (
	"BE-SISFO-KLINIK/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectionDatabase() (*gorm.DB, error) {
	dsn := "root:fanzru@tcp(0.0.0.0:1000)/db_klinik?charset=utf8mb4&parseTime=True&loc=Local"
	//	dsn := "root:@tcp(127.0.0.1:3306)/db_klinik?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if there is an error opening the connection, handle it
	if err != nil {
		return nil, err
	}

	return db, nil
}
func MigrationDB() (*gorm.DB, error) {

	fmt.Println("Connecting to Database ....")
	db, err := ConnectionDatabase()
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&models.Pasien{})
	if err != nil {
		fmt.Println("Migration table pasiens failed")
		return nil, err
	}
	err = db.AutoMigrate(&models.Pemeriksaan{})
	if err != nil {
		fmt.Println("Migration table pemeriksaans failed")
		return nil, err
	}
	err = db.AutoMigrate(&models.Obat{})
	if err != nil {
		fmt.Println("Migration table obats failed")
		return nil, err
	}
	err = db.AutoMigrate(&models.StokObat{})
	if err != nil {
		fmt.Println("Migration table stokobats failed")
		return nil, err
	}
	err = db.AutoMigrate(&models.Obat{})
	if err != nil {
		fmt.Println("Migration table obats failed")
		return nil, err
	}
	err = db.AutoMigrate(&models.TransaksiObat{})
	if err != nil {
		fmt.Println("Migration table transaksi_obats failed")
		return nil, err
	}
	// err = db.AutoMigrate(&models.Transaksi{})
	// if err != nil {
	// 	fmt.Println("Migration table transaksi failed")
	// 	return nil, err
	// }

	// add index in database 'nama_pasien' ALTER TABLE `table_name` ADD FULLTEXT(`column_name`);
	// err = db.Raw("ALTER TABLE `pasiens` ADD FULLTEXT(`nama_pasien`);").Error
	// if err != nil {
	// 	fmt.Println("Add FULL TEXT IN PASIENS")
	// 	return nil, err
	// }
	fmt.Println("Database Connected")
	return db, nil
}
