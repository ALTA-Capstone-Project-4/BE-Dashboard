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

	if data.ID == id {
		tx := repo.db.Model(&Gudang{}).Updates(dataModel)
		if tx.Error != nil {
			return -1, tx.Error
		}
		return 1, nil
	}
	return 1, nil
}

func (repo *gudangData) SelectAllGudang() ([]gudang.Core, error) {
	var dataGudang []Gudang

	tx := repo.db.Model(&Gudang{}).Find(&dataGudang)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return toCoreList(dataGudang), nil
}
