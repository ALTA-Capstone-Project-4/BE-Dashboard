package data

import (
	"warehouse/features/gudang"

	"gorm.io/gorm"
)

type gudangData struct {
	db *gorm.DB
}

func New(db *gorm.DB) gudang.DataInterface {
	return &gudangData{
		db: db,
	}
}

func (repo *gudangData) UpdateGudang(id int, data gudang.Core) (int, error) {
	dataModel := fromCore(data)
	tx := repo.db.Model(&Gudang{}).Where("user_id = ?", id).Updates(dataModel)
	if tx.Error != nil {
		return -1, tx.Error
	}
	return 1, nil
}
