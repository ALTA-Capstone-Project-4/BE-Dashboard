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

func (repo *gudangData) SelectAllLahan(offset int) ([]gudang.Lahan, error) {
	var dataGudang []Lahan

	tx_hargaMin := repo.db.Model(&Lahan{}).Offset(offset).Limit(8).Select("gudang_id, MIN(harga) as harga").Group("gudang_id").Find(&dataGudang)

	if tx_hargaMin.Error != nil {
		return nil, tx_hargaMin.Error
	}

	for key := range dataGudang {
		tx := repo.db.Model(&Lahan{}).Offset(offset).Limit(8).Where("harga = ? AND gudang_id = ?", dataGudang[key].Harga, dataGudang[key].GudangID).Find(&dataGudang[key])
		if tx.Error != nil {
			return nil, tx.Error
		}
	}

	return toLahanList(dataGudang), nil
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
	dataLahan := toLahanList(gudangData.Lahan)
	detailGudang.Lahan = dataLahan

	for key := range detailGudang.Lahan {
		detailGudang.Lahan[key].Alamat = gudangData.Location
		detailGudang.Lahan[key].Nama_Gudang = gudangData.Name
	}

	return detailGudang, nil
}
