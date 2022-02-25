package config

import (
	"BE-SISFO-KLINIK/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectionDatabase() (*gorm.DB, error) {
	dsn := "root:@tcp(127.0.0.1:3306)/db_klinik?charset=utf8mb4&parseTime=True&loc=Local"
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
		fmt.Println("Migration failed")
		return nil, err
	}
	err = db.AutoMigrate(&models.Pemeriksaan{})
	if err != nil {
		fmt.Println("Migration failed")
		return nil, err
	}

	fmt.Println("Database Connected")
	return db, nil
}
