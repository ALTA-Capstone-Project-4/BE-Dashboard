package data

import (
	"warehouse/features/lahan"

	"gorm.io/gorm"
)

type lahanData struct {
	db *gorm.DB
}

func New(db *gorm.DB) lahan.DataInterface {
	return &lahanData{
		db: db,
	}
}
