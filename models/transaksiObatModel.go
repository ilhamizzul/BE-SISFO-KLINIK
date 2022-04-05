package models

import "time"

type (
	TransaksiObat struct {
		Id            int64     `gorm:"AUTO_INCREMENT;PRIMARY_KEY;not null"`
		RincianObat   string    `gorm:"type:string;not null" validate:"required"`
		Jumlah        int64     `gorm:"type:integer;not null;" validate:"required"`
		Harga         int64     `gorm:"type:integer;not null;" validate:"required"`
		IdObat        int64     `gorm:"foreignKey:IdObat;references:Obats;not null" validate:"required"`
		IdPemeriksaan int64     `gorm:"foreignKey:IdPemeriksaan;references:Pemeriksaans;not null" validate:"required"`
		DeleteStatus  bool      `gorm:"type:bool;not null;default:0"`
		CreatedAt     time.Time `gorm:"type:datetime;not null"`
		UpdatedAt     time.Time `gorm:"type:datetime;"`
	}
)
