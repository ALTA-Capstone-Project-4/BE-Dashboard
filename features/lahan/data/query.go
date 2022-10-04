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

func (repo *lahanData) SelectDetailLahan(id int, role string) (lahan.Core, error) {
	var data Lahan
	tx := repo.db.Where("id = ?", id).Find(&data)
	if tx.Error != nil {
		return lahan.Core{}, tx.Error
	}

	return data.toCore(), nil
}

func (repo *lahanData) UpdateLahan(id int, token int, data lahan.Core) (int, error) {
	var gudangModel modelGudang.Gudang
	repo.db.Where("user_id = ?", token).Find(&gudangModel)
	data.GudangID = gudangModel.ID

	dataModel := fromCore(data)
	tx := repo.db.Model(&Lahan{}).Where("id = ? AND gudang_id = ?", id, data.GudangID).Updates(&dataModel)
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected < 1 {
		return -1, errors.New("Unauthorized")
	}
	return 1, nil
}

func (repo *lahanData) DeleteData(id int, token int, data lahan.Core) (int, error) {
	var gudangModel modelGudang.Gudang
	repo.db.Where("user_id = ?", token).Find(&gudangModel)
	data.GudangID = gudangModel.ID

	var deleteData Lahan
	tx := repo.db.Where("id = ? AND gudang_id = ?", id, data.GudangID).Delete(&deleteData)
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected < 1 {
		return -1, errors.New("Unauthorized")
	}

	return 1, nil
}

func (repo *lahanData) SelectLahanClient(token int) ([]lahan.Core, error) {
	// var checkoutModel modelGudang.Gudang
	var data []Lahan
	tx := repo.db.Model(&Lahan{}).Where("status = ?", "verified").Find(&data)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return toCoreList(data), nil
}
