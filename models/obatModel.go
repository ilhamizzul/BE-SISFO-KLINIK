package models

import (
	"time"
)

type (
	Obat struct {
		Id           int64     `gorm:"AUTO_INCREMENT;PRIMARY_KEY;not null"`
		Kode         string    `gorm:"type:string;unique;not null;index:;class:FULLTEXT"  validate:"required"`
		Nama         string    `gorm:"type:string;not null" validate:"required"`
		HargaJual    int64     `gorm:"type:integer;not null" validate:"required"`
		DeleteStatus bool      `gorm:"type:bool;not null"`
		CreatedAt    time.Time `gorm:"type:datetime;not null"`
		UpdatedAt    time.Time `gorm:"type:datetime;"`
	}
)
