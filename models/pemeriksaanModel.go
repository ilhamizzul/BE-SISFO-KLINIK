package models

import (
	"time"
)

type (
	Pemeriksaan struct {
		IdPemeriksaan      int64     `gorm:"AUTO_INCREMENT;PRIMARY_KEY;not null"`
		TanggalPemeriksaan time.Time `gorm:"type:datetime;not null"`
		HasilPemeriksaan   string    `gorm:"string;not null"`
		Diagnosis          string    `gorm:"string;not null"`
		Terapi             string    `gorm:"string;not null"`
		CatatanDokter      string    `gorm:"string;not null"`
		StatusTransaksi    string    `gorm:"type:enum('sudah', 'belum');default:'belum'"`
		IdPasien           int64     `gorm:"foreignKey:IdPasien;references:Pasien"`
		DeleteStatus       bool      `gorm:"type:bool;not null"`
		CreatedAt          time.Time `gorm:"type:datetime;not null"`
		UpdatedAt          time.Time `gorm:"type:datetime;"`
	}
)
