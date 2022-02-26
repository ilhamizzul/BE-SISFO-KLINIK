package models

import (
	"time"
)

type (
	Pasien struct {
		IdPemeriksaan      int64     `gorm:"AUTO_INCREMENT;PRIMARY_KEY;not null"`
		NamaPasien         string    `gorm:"type:string;not null" validate:"required"`
		Alamat             string    `gorm:"type:string;unique;not null" validate:"required"`
		TempatLahir        string    `gorm:"type:string;not null" validate:"required"`
		TanggalLahir       time.Time `gorm:"type:datetime;not null" validate:"required"`
		NamaKepalaKeluarga string    `gorm:"type:string;not null" validate:"required"`
		DeleteStatus       bool      `gorm:"type:bool;not null"`
		CreatedAt          time.Time `gorm:"type:datetime;not null"`
		UpdatedAt          time.Time `gorm:"type:datetime;"`
	}
)
