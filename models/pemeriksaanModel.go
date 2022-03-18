package models

import (
	"time"
)

type (
	Pemeriksaan struct {
		Id                 int64     `gorm:"AUTO_INCREMENT;PRIMARY_KEY;not null"`
		TanggalPemeriksaan time.Time `gorm:"type:datetime;not null" validate:"required"`
		HasilPemeriksaan   string    `gorm:"type:string;not null" validate:"required"`
		Diagnosis          string    `gorm:"type:string;not null" validate:"required"`
		Terapi             string    `gorm:"type:string;not null" validate:"required"`
		CatatanDokter      string    `gorm:"type:string;not null" validate:"required"`
		StatusTransaksi    string    `gorm:"type:enum('sudah', 'belum');default:'belum'"`
		IdPasien           int64     `gorm:"foreignKey:IdPasien;references:Pasiens" validate:"required"`
		DeleteStatus       bool      `gorm:"type:bool;not null"`
		CreatedAt          time.Time `gorm:"type:datetime;not null"`
		UpdatedAt          time.Time `gorm:"type:datetime;"`
	}
)
