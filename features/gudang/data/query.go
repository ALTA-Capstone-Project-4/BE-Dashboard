package data

import (
	"errors"
	"warehouse/features/gudang"
	userModel "warehouse/features/user/data"

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

	tx := repo.db.Where("user_id = ?", id).Model(&Gudang{}).Updates(dataModel)
	if tx.Error != nil {
		return -1, tx.Error
	}
	return int(tx.RowsAffected), nil

}

func (repo *gudangData) SelectAllGudang(offset int) ([]gudang.Lahan, error) {
	var dataGudang []Lahan

	tx := repo.db.Model(&Lahan{}).Offset(offset).Limit(8).Select("panjang, lebar, MIN(harga) as harga, foto_lahan, gudang_id").Group("gudang_id").Find(&dataGudang)

	if tx.Error != nil {
		return nil, tx.Error
	}
	return toLahanCoreList(dataGudang), nil
}

func (repo *gudangData) CreatGudang(data gudang.Core) (int, error) {
	var userData userModel.User
	dataModel := fromCore(data)

	tx_user := repo.db.Where("id = ?", data.UserID).Find(&userData)

	if tx_user.Error != nil {
		return 0, tx_user.Error
	}

	if userData.Status != "verified" {
		return -1, errors.New("your account unverified")
	}

	var tempGudangData Gudang
	tx_gudang := repo.db.Where("user_id = ?", userData.ID).Find(&tempGudangData)

	if tx_gudang.Error != nil {
		return 0, tx_user.Error
	}

	if tempGudangData.ID != 0 {
		return -1, errors.New("can't add gudang, already create one")
	}

	tx := repo.db.Create(&dataModel)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return int(tx.RowsAffected), nil
}

func (repo *gudangData) SelectGudangByID(gudang_id int) (gudang.Core, error) {
	var gudangData Gudang

	tx := repo.db.Where("id = ?", gudang_id).Preload("Lahan").Find(&gudangData)
	if tx.Error != nil {
		return gudang.Core{}, tx.Error
	}

	detailGudang := gudangData.toCore()
	dataLahan := ToLahanList(gudangData.Lahan)
	detailGudang.Lahan = dataLahan

	return detailGudang, nil
}
