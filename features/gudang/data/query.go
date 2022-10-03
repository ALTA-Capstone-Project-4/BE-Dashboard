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

	if data.ID == id {
		tx := repo.db.Model(&Gudang{}).Updates(dataModel)
		if tx.Error != nil {
			return -1, tx.Error
		}
		return 1, nil
	}
	return 1, nil
}

func (repo *gudangData) SelectAllGudang() ([]gudang.Lahan, error) {
	var dataGudang []Lahan

	tx := repo.db.Model(&Lahan{}).Select("luas, MIN(harga) as harga, foto_lahan, gudang_id").Group("gudang_id").Find(&dataGudang)
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
	tx := repo.db.Create(&dataModel)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return int(tx.RowsAffected), nil
}
