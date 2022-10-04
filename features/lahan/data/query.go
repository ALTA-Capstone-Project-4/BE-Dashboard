package data

import (
	"errors"
	modelGudang "warehouse/features/gudang/data"
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

func (repo *lahanData) CreateLahan(data lahan.Core, user_id int) (int, error) {

	var gudangModel modelGudang.Gudang
	repo.db.Where("user_id = ?", user_id).Find(&gudangModel)
	data.GudangID = gudangModel.ID

	dataModel := fromCore(data)

	tx := repo.db.Create(&dataModel)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return int(tx.RowsAffected), nil
}

func (repo *lahanData) UpdateLahan(id int, data lahan.Core) (int, error) {
	dataModel := fromCore(data)
	tx := repo.db.Model(&Lahan{}).Where("id = ?", id).Updates(dataModel)
	if tx.Error != nil {
		return -1, tx.Error
	}
	return 1, nil
}

func (repo *lahanData) SelectDetailLahan(id int, role string) (lahan.Core, error) {

	if role == "mitra" {

		var data Lahan
		tx := repo.db.Where("id = ?", id).Preload("Checkout").Find(&data)
		if tx.Error != nil {
			return lahan.Core{}, tx.Error
		}

		return data.toCore(), nil

	} else {

		var data Lahan
		tx := repo.db.Where("id = ?", id).Find(&data)
		if tx.Error != nil {
			return lahan.Core{}, tx.Error
		}

		return data.toCore(), nil
	}
}

func (repo *lahanData) DeleteData(token int, id int) (int, error) {
	var data Gudang
	tx := repo.db.Where("user_id = ?", token).Delete(&data)
	if tx.Error != nil {
		return -1, errors.New("Unauthorized")
	}

	var deleteData Lahan
	txLahan := repo.db.Where("id = ?", id).Delete(&deleteData)
	if txLahan.Error != nil {
		return -1, txLahan.Error
	}

	return 1, nil

}
