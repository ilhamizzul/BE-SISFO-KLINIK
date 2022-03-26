package models

import (
	"time"
)

type (
	StokObat struct {
		Id           int64     `gorm:"AUTO_INCREMENT;PRIMARY_KEY;not null"`
		DeleteStatus bool      `gorm:"type:bool;not null"`
		Masuk        int64     `gorm:"type:integer;not null;default:0" validate:"required"`
		Keluar       int64     `gorm:"type:integer;not null;default:0" validate:"required"`
		Sisa         int64     `gorm:"type:integer;not null;default:0" validate:"required"`
		Date         time.Time `gorm:"type:datetime;not null"`
		IdObat       int64     `gorm:"foreignKey:IdObat;references:Obats;not null" validate:"required"`
		CreatedAt    time.Time `gorm:"type:datetime;not null"`
		UpdatedAt    time.Time `gorm:"type:datetime;"`
	}
)
